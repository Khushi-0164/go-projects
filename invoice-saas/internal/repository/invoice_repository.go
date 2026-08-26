package repository

import (
	"invoice-saas/internal/models"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}
func (r *InvoiceRepository) CreateWithLineItems(invoice *models.Invoice, lineItems []models.InvoiceLineItem) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var total int64
		for _, item := range lineItems {
			total += int64(item.Quantity) * item.UnitPriceCents
		}
		invoice.TotalCents = total
		for i := range lineItems {
			lineItems[i].InvoiceID = invoice.ID
		}
		if err := tx.Create(&lineItems).Error; err != nil {
			return err
		}

		invoice.LineItems = lineItems
		return nil
	})
}

func (r *InvoiceRepository) FindAll(orgID uint) ([]models.Invoice, error) {
	var invoices []models.Invoice

	err := r.DB.Where("organization_id = ?", orgID).Order("created_at desc").Find(&invoices).Error
	return invoices, err
}

func (r *InvoiceRepository) FindByID(orgID, invoiceID uint) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.DB.
		Where("organization_id = ? AND id = ?", orgID, invoiceID).
		Preload("LineItems").
		First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *InvoiceRepository) UpdateStatus(orgID, invoiceID uint, status models.InvoiceStatus) error {
	result := r.DB.Model(&models.Invoice{}).
		Where("organization_id = ? AND id = ?", orgID, invoiceID).
		Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
