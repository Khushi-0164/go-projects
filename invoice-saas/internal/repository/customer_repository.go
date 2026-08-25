package repository

import (
	"invoice-saas/internal/models"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	return r.DB.Create(customer).Error
}

func (r *CustomerRepository) FindAll(orgID uint) ([]models.Customer, error) {
	var customers []models.Customer
	err := r.DB.Where("organization_id=?", orgID).Find(&customers).Error
	return customers, err
}

func (r *CustomerRepository) FindByID(orgID, customerID uint) (*models.Customer, error) {
	var customer models.Customer
	err := r.DB.Where("organization_id=? AND id=?", orgID, customerID).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
