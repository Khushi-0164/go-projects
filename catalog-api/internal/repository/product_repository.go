package repository

import (
	"catalog-api/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

const cacheTTL = 5 * time.Minute

type ProductRepository struct {
	DB    *gorm.DB
	Cache *redis.Client
}

func NewProductRepository(db *gorm.DB, cache *redis.Client) *ProductRepository {
	return &ProductRepository{DB: db, Cache: cache}
}

func cacheKey(id uint) string {
	return fmt.Sprintf("product:%d", id)
}

func (r *ProductRepository) FindByID(ctx context.Context, id uint) (*models.Product, error) {
	key := cacheKey(id)

	cached, err := r.Cache.Get(ctx, key).Result()
	if err == nil {
		var product models.Product
		if jsonErr := json.Unmarshal([]byte(cached), &product); jsonErr == nil {
			return &product, nil
		}
	}
	time.Sleep(300 * time.Millisecond)
	var product models.Product
	if dbErr := r.DB.First(&product, id).Error; dbErr != nil {
		return nil, dbErr
	}
	if encoded, jsonErr := json.Marshal(product); jsonErr == nil {
		r.Cache.Set(ctx, key, encoded, cacheTTL)
	}
	return &product, nil
}

func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	if err := r.DB.Save(product).Error; err != nil {
		return err
	}
	r.Cache.Del(ctx, cacheKey(product.ID))
	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id uint) error {
	if err := r.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}
	r.Cache.Del(ctx, cacheKey(id))
	return nil
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}
