package service

import (
	"invoice-saas/internal/models"
)

type CustomerRepository interface {
	Create(customer *models.Customer) error
	FindAll(orgID uint) ([]models.Customer, error)
	FindByID(orgID, customerID uint) (*models.Customer, error)
}

type CustomerService struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) CreateCustomer(orgID uint, name, email string) (*models.Customer, error) {
	customer := &models.Customer{
		OrganizationID: orgID,
		Name:           name,
		Email:          email,
	}
	if err := s.repo.Create(customer); err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *CustomerService) ListCustomers(orgID uint) ([]models.Customer, error) {
	return s.repo.FindAll(orgID)
}

func (s *CustomerService) GetCustomer(orgID, customerID uint) (*models.Customer, error) {
	return s.repo.FindByID(orgID, customerID)
}
