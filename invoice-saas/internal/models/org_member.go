package models

import "time"

type OrgRole string

const (
	OrgRoleOwner  OrgRole = "owner"
	OrgRoleAdmin  OrgRole = "admin"
	OrgRoleMember OrgRole = "member"
)

type OrgMember struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	OrganizationID uint      `gorm:"not null;uniqueIndex:idx_org_user" json:"organization_id"`
	UserID         uint      `gorm:"not null;uniqueIndex:idx_org_user" json:"user_id"`
	Role           OrgRole   `gorm:"type:varchar(20);not null" json:"role"`
	CreatedAt      time.Time `json:"created_at"`
}
