package service

import (
	"errors"

	"task-manager/internal/models"
	"task-manager/internal/repository"
)

var (
	ErrNotAuthorized = errors.New("not authorized to perform this action")
	ErrUserNotFound  = errors.New("user not found")
	ErrTaskNotFound  = errors.New("task not found")
)

// ProjectService contains the business rules for projects and membership.
type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// CreateProject creates a project and makes the creator its owner.
func (s *ProjectService) CreateProject(name, description string, ownerID uint) (*models.Project, error) {
	project := &models.Project{
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}
	if err := s.repo.CreateProjectWithOwner(project, ownerID); err != nil {
		return nil, err
	}
	return project, nil
}

// ListMyProjects returns every project a user belongs to.
func (s *ProjectService) ListMyProjects(userID uint) ([]models.Project, error) {
	return s.repo.FindProjectsForUser(userID)
}

// AddMember adds a user to a project by email, enforcing that only
// owners/admins of that specific project can invite others.
func (s *ProjectService) AddMember(projectID, requesterID uint, targetEmail string, role models.ProjectRole) (*models.ProjectMember, error) {
	requesterRole, isMember := s.repo.FindMemberRole(projectID, requesterID)
	if !isMember || (requesterRole != models.ProjectRoleOwner && requesterRole != models.ProjectRoleAdmin) {
		return nil, ErrNotAuthorized
	}

	targetUser, err := s.repo.FindUserByEmail(targetEmail)
	if err != nil {
		return nil, ErrUserNotFound
	}

	member := &models.ProjectMember{
		ProjectID: projectID,
		UserID:    targetUser.ID,
		Role:      role,
	}
	if err := s.repo.AddMember(member); err != nil {
		return nil, err
	}
	return member, nil

}
