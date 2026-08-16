package service

import (
	"catalog-api/internal/models"
	"catalog-api/internal/repository"
	"context"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(ctx context.Context, id uint) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context) ([]models.Product, error) {
	return s.repo.FindAll(ctx)
}

func (s *ProductService) CreateProduct(ctx context.Context, name, description string, price float64) (*models.Product, error) {
	product := &models.Product{Name: name, Description: description, Price: price}
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id uint, name, description string, price float64) (*models.Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	product.Name = name
	product.Description = description
	product.Price = price

	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
