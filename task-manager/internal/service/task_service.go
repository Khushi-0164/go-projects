package service

import (
	"task-manager/internal/models"
	"task-manager/internal/repository"
)

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(projectID, requesterID uint, title, description string, assigneeID *uint) (*models.Task, error) {
	if _, isMember := s.repo.FindMemberRole(projectID, requesterID); !isMember {
		return nil, ErrNotAuthorized
	}

	task := &models.Task{
		ProjectID:   projectID,
		Title:       title,
		Description: description,
		Status:      models.StatusTodo,
		AssigneeID:  assigneeID,
	}
	if err := s.repo.Create(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) ListTasks(projectID, requesterID uint) ([]models.Task, error) {
	if _, isMember := s.repo.FindMemberRole(projectID, requesterID); !isMember {
		return nil, ErrNotAuthorized
	}
	return s.repo.FindByProject(projectID)
}

func (s *TaskService) UpdateTaskStatus(taskID, requesterID uint, status models.TaskStatus) (*models.Task, error) {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return nil, ErrTaskNotFound
	}

	if _, isMember := s.repo.FindMemberRole(task.ProjectID, requesterID); !isMember {
		return nil, ErrNotAuthorized
	}

	task.Status = status
	if err := s.repo.Save(task); err != nil {
		return nil, err
	}
	return task, nil
}
