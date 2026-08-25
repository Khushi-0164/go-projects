package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	OrganizationID uint           `gorm:"not null;index" json:"organization_id"`
	Name           string         `gorm:"not null" json:"name"`
	Email          string         `gorm:"not null" json:"email"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
