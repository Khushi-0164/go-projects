package models

import (
	"time"

	"gorm.io/gorm"
)

type InvoiceStatus string

const (
	InvoiceStatusDraft   InvoiceStatus = "draft"
	InvoiceStatusSent    InvoiceStatus = "sent"
	InvoiceStatusPaid    InvoiceStatus = "paid"
	InvoiceStatusOverdue InvoiceStatus = "overdue"
	InvoiceStatusVoid    InvoiceStatus = "void"
)

type Invoice struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	OrganizationID uint           `gorm:"not null;index" json:"organization_id"`
	CustomerID     uint           `gorm:"not null;index" json:"customer_id"`
	Status         InvoiceStatus  `gorm:"type:varchar(20);default:'draft'" json:"status"`
	TotalCents     int64          `gorm:"not null;default:0" json:"total_cents"`
	DueDate        *time.Time     `json:"due_date"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	LineItems []InvoiceLineItem `gorm:"foreignKey:InvoiceID" json:"line_items,omitempty"`
}
