package models

import "time"

type GroupMember struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	GroupID   uint      `gorm:"not null;uniqueIndex:idx_group_user" json:"group_id"`
	UserID    uint      `gorm:"not null;uniqueIndex:idx_group_user" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
