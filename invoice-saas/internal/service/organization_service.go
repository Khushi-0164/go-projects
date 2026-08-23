package service

import (
	"errors"
	"invoice-saas/internal/models"
)

var (
	ErrNotAuthorized = errors.New("not authorized to perform this action")
	ErrUserNotFound  = errors.New("user not found")
)

type OrganizationRepository interface {
	CreateWithOwner(org *models.Organization, ownerID uint) error
	FindOrgsForUser(userID uint) ([]models.Organization, error)
	FindMemberRole(orgID, userID uint) (models.OrgRole, bool)
	FindUserByEmail(email string) (*models.User, error)
	AddMember(member *models.OrgMember) error
}

type OrganizationService struct {
	repo OrganizationRepository
}

func NewOrganizationService(repo OrganizationRepository) *OrganizationService {
	return &OrganizationService{repo: repo}
}

func (s *OrganizationService) CreateOrganization(name string, ownerID uint) (*models.Organization, error) {
	org := &models.Organization{Name: name}
	if err := s.repo.CreateWithOwner(org, ownerID); err != nil {
		return nil, err
	}
	return org, nil
}

func (s *OrganizationService) ListMyOrganizations(userID uint) ([]models.Organization, error) {
	return s.repo.FindOrgsForUser(userID)
}
func (s *OrganizationService) AddMember(orgID, requesterID uint, targetEmail string, role models.OrgRole) (*models.OrgMember, error) {
	requesterRole, isMember := s.repo.FindMemberRole(orgID, requesterID)
	if !isMember || (requesterRole != models.OrgRoleOwner && requesterRole != models.OrgRoleAdmin) {
		return nil, ErrNotAuthorized
	}

	targetUser, err := s.repo.FindUserByEmail(targetEmail)
	if err != nil {
		return nil, ErrUserNotFound
	}

	member := &models.OrgMember{
		OrganizationID: orgID,
		UserID:         targetUser.ID,
		Role:           role,
	}
	if err := s.repo.AddMember(member); err != nil {
		return nil, err
	}
	return member, nil
}
