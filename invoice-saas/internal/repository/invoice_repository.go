package repository

import (
	"context"
	"invoice-saas/internal/models"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	DB      *gorm.DB
	Summary *CachedSummaryRepository
}

func NewInvoiceRepository(db *gorm.DB, summary *CachedSummaryRepository) *InvoiceRepository {
	return &InvoiceRepository{DB: db, Summary: summary}
}
func (r *InvoiceRepository) CreateWithLineItems(invoice *models.Invoice, lineItems []models.InvoiceLineItem) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		var total int64
		for _, item := range lineItems {
			total += int64(item.Quantity) * item.UnitPriceCents
		}
		invoice.TotalCents = total

		if err := tx.Create(invoice).Error; err != nil {
			return err
		}

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
func (r *InvoiceRepository) FindAll(orgID uint, page, limit int, status string) ([]models.Invoice, int64, error) {
	var invoices []models.Invoice
	var total int64

	query := r.DB.Model(&models.Invoice{}).Where("organization_id = ?", orgID)
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Order("created_at desc").Limit(limit).Offset(offset).Find(&invoices).Error

	return invoices, total, err
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

	r.Summary.InvalidateSummary(context.Background(), orgID)
	return nil
}
