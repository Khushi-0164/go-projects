package repository

import (
	"booking-api/internal/models"

	"gorm.io/gorm"
)

type ResourceRepository struct {
	DB *gorm.DB
}

func NewResourceRepository(db *gorm.DB) *ResourceRepository {
	return &ResourceRepository{DB: db}
}

func (r *ResourceRepository) Create(resource *models.Resource) error {
	return r.DB.Create(resource).Error
}

func (r *ResourceRepository) FindAll() ([]models.Resource, error) {
	var resources []models.Resource
	err := r.DB.Find(&resources).Error
	return resources, err
}

func (r *ResourceRepository) FindByID(id uint) (*models.Resource, error) {
	var resource models.Resource
	if err := r.DB.First(&resource, id).Error; err != nil {
		return nil, err
	}
	return &resource, nil
}
