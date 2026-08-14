package models

import "time"

type Booking struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ResourceID uint      `gorm:"not null;index" json:"resource_id"`
	UserID     uint      `gorm:"not null;index" json:"user_id"`
	StartTime  time.Time `gorm:"not null" json:"start_time"`
	EndTime    time.Time `gorm:"not null" json:"end_time"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `gorm:"index" json:"-"`
}
