package models

import "time"

type InvoiceLineItem struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	InvoiceID      uint      `gorm:"not null;index" json:"invoice_id"`
	Description    string    `gorm:"not null" json:"description"`
	Quantity       int       `gorm:"not null;default:1" json:"quantity"`
	UnitPriceCents int64     `gorm:"not null" json:"unit_price_cents"`
	CreatedAt      time.Time `json:"created_at"`
}
