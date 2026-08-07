package repository

import (
	"task-manager/internal/models"

	"gorm.io/gorm"
)

// ProjectRepository handles all direct database access for projects and memberships.
type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{DB: db}
}

func (r *ProjectRepository) CreateProjectWithOwner(project *models.Project, ownerID uint) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(project).Error; err != nil {
			return err
		}
		member := models.ProjectMember{
			ProjectID: project.ID,
			UserID:    ownerID,
			Role:      models.ProjectRoleOwner,
		}
		return tx.Create(&member).Error
	})
}

func (r *ProjectRepository) FindProjectsForUser(userID uint) ([]models.Project, error) {
	var projects []models.Project
	err := r.DB.
		Joins("JOIN project_members ON project_members.project_id = projects.id").
		Where("project_members.user_id = ?", userID).
		Find(&projects).Error
	return projects, err
}

func (r *ProjectRepository) FindMemberRole(projectID, userID uint) (models.ProjectRole, bool) {
	var member models.ProjectMember
	if err := r.DB.Where("project_id = ? AND user_id = ?", projectID, userID).First(&member).Error; err != nil {
		return "", false
	}
	return member.Role, true
}

func (r *ProjectRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *ProjectRepository) AddMember(member *models.ProjectMember) error {
	return r.DB.Create(member).Error
}
