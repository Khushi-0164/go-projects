package models

import "time"

type ProjectRole string

const (
	ProjectRoleOwner  ProjectRole = "owner"
	ProjectRoleAdmin  ProjectRole = "admin"
	ProjectRoleMember ProjectRole = "member"
)

type ProjectMember struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	ProjectID uint        `gorm:"not null;uniqueIndex:idx_project_user" json:"project_id"`
	UserID    uint        `gorm:"not null;uniqueIndex:idx_project_user" json:"user_id"`
	Role      ProjectRole `gorm:"type:varchar(20);not null" json:"role"`
	CreatedAt time.Time   `json:"created_at"`

	User User `gorm:"foreignKey:UserID" json:"user"`
}
