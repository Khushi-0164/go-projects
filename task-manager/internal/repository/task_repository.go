package repository

import (
	"task-manager/internal/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) Create(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) FindByProject(projectID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.DB.Where("project_id = ?", projectID).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) FindByID(taskID uint) (*models.Task, error) {
	var task models.Task
	if err := r.DB.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Save(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) FindMemberRole(projectID, userID uint) (models.ProjectRole, bool) {
	var member models.ProjectMember
	if err := r.DB.Where("project_id = ? AND user_id = ?", projectID, userID).First(&member).Error; err != nil {
		return "", false
	}
	return member.Role, true
}
