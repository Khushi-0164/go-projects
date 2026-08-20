package repository

import (
	"newsletter-api/internal/models"

	"gorm.io/gorm"
)

type SubscriberRepository struct {
	DB *gorm.DB
}

func NewSubscriberRepository(db *gorm.DB) *SubscriberRepository {
	return &SubscriberRepository{DB: db}
}

func (r *SubscriberRepository) Create(subscriber *models.Subscriber) error {
	return r.DB.Create(subscriber).Error
}

func (r *SubscriberRepository) ExistsByEmail(email string) bool {
	var count int64
	r.DB.Model(&models.Subscriber{}).Where("email = ?", email).Count(&count)
	return count > 0
}
