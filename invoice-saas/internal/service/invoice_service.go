package service

import (
	"invoice-saas/internal/models"
)

type InvoiceRepository interface {
	CreateWithLineItems(invoice *models.Invoice, lineItems []models.InvoiceLineItem) error
	FindAll(orgID uint, page, limit int, status string) ([]models.Invoice, int64, error)
	FindByID(orgID, invoiceID uint) (*models.Invoice, error)
	UpdateStatus(orgID, invoiceID uint, status models.InvoiceStatus) error
}

type InvoiceService struct {
	repo InvoiceRepository
}

func NewInvoiceService(repo InvoiceRepository) *InvoiceService {
	return &InvoiceService{repo: repo}
}

type LineItemInput struct {
	Description    string
	Quantity       int
	UnitPriceCents int64
}

func (s *InvoiceService) CreateInvoice(orgID, customerID uint, items []LineItemInput) (*models.Invoice, error) {
	invoice := &models.Invoice{
		OrganizationID: orgID,
		CustomerID:     customerID,
		Status:         models.InvoiceStatusDraft,
	}

	lineItems := make([]models.InvoiceLineItem, len(items))
	for i, item := range items {
		lineItems[i] = models.InvoiceLineItem{
			Description:    item.Description,
			Quantity:       item.Quantity,
			UnitPriceCents: item.UnitPriceCents,
		}
	}

	if err := s.repo.CreateWithLineItems(invoice, lineItems); err != nil {
		return nil, err
	}
	return invoice, nil
}

func (s *InvoiceService) ListInvoices(orgID uint, page, limit int, status string) ([]models.Invoice, int64, error) {
	return s.repo.FindAll(orgID, page, limit, status)
}

func (s *InvoiceService) GetInvoice(orgID, invoiceID uint) (*models.Invoice, error) {
	return s.repo.FindByID(orgID, invoiceID)
}

func (s *InvoiceService) UpdateStatus(orgID, invoiceID uint, status models.InvoiceStatus) error {
	return s.repo.UpdateStatus(orgID, invoiceID, status)
}
