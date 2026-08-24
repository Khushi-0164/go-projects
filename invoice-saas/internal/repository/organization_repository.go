package repository

import (
	"invoice-saas/internal/models"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	DB *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{DB: db}
}

func (r *OrganizationRepository) CreateWithOwner(org *models.Organization, ownerID uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(org).Error; err != nil {
			return err
		}
		member := models.OrgMember{
			OrganizationID: org.ID,
			UserID:         ownerID,
			Role:           models.OrgRoleOwner,
		}
		return tx.Create(&member).Error
	})
}
func (r *OrganizationRepository) FindOrgsForUser(userID uint) ([]models.Organization, error) {
	var orgs []models.Organization
	err := r.DB.
		Joins("JOIN org_members ON org_members.organization_id = organizations.id").
		Where("org_members.user_id = ?", userID).
		Find(&orgs).Error
	return orgs, err
}
func (r *OrganizationRepository) FindMemberRole(orgID, userID uint) (models.OrgRole, bool) {
	var member models.OrgMember
	if err := r.DB.Where("organization_id = ? AND user_id = ?", orgID, userID).First(&member).Error; err != nil {
		return "", false
	}
	return member.Role, true
}
func (r *OrganizationRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *OrganizationRepository) AddMember(member *models.OrgMember) error {
	return r.DB.Create(member).Error
}
