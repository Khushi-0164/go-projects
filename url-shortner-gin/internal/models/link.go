package models

import (
	"time"

	"gorm.io/gorm"
)

// Link represents a shortened URL owned by a user.
type Link struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	ShortCode   string         `gorm:"uniqueIndex;size:16;not null" json:"short_code"`
	OriginalURL string         `gorm:"not null" json:"original_url"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	Clicks      uint           `gorm:"default:0" json:"clicks"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
