package service

import (
	"testing"

	"task-manager/internal/models"
	"task-manager/internal/repository"
)

func TestCreateTask_RejectsNonMember(t *testing.T) {
	db := setupTestDB(t)
	projectRepo := repository.NewProjectRepository(db)
	projectSvc := NewProjectService(projectRepo)
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := NewTaskService(taskRepo)

	owner := createTestUser(t, db, "owner3@example.com")
	outsider := createTestUser(t, db, "outsider@example.com")

	project, err := projectSvc.CreateProject("Project X", "", owner.ID)
	if err != nil {
		t.Fatalf("setup failed: could not create project: %v", err)
	}

	// outsider is NOT a member of this project — should be rejected
	_, err = taskSvc.CreateTask(project.ID, outsider.ID, "Sneaky task", "", nil)
	if err == nil {
		t.Fatalf("expected an error when a non-member tries to create a task, got nil")
	}
	if err != ErrNotAuthorized {
		t.Errorf("expected ErrNotAuthorized, got: %v", err)
	}
}

func TestCreateTask_SucceedsForMember(t *testing.T) {
	db := setupTestDB(t)
	projectRepo := repository.NewProjectRepository(db)
	projectSvc := NewProjectService(projectRepo)
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := NewTaskService(taskRepo)

	owner := createTestUser(t, db, "owner4@example.com")

	project, err := projectSvc.CreateProject("Project Y", "", owner.ID)
	if err != nil {
		t.Fatalf("setup failed: could not create project: %v", err)
	}

	task, err := taskSvc.CreateTask(project.ID, owner.ID, "Write docs", "some description", nil)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if task.Status != models.StatusTodo {
		t.Errorf("expected new task status to be %q, got %q", models.StatusTodo, task.Status)
	}
	if task.AssigneeID != nil {
		t.Errorf("expected AssigneeID to be nil for an unassigned task, got %v", task.AssigneeID)
	}
}

func TestUpdateTaskStatus_ChangesStatus(t *testing.T) {
	db := setupTestDB(t)
	projectRepo := repository.NewProjectRepository(db)
	projectSvc := NewProjectService(projectRepo)
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := NewTaskService(taskRepo)

	owner := createTestUser(t, db, "owner5@example.com")
	project, _ := projectSvc.CreateProject("Project Z", "", owner.ID)
	task, _ := taskSvc.CreateTask(project.ID, owner.ID, "Ship feature", "", nil)

	updated, err := taskSvc.UpdateTaskStatus(task.ID, owner.ID, models.StatusInProgress)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if updated.Status != models.StatusInProgress {
		t.Errorf("expected status %q, got %q", models.StatusInProgress, updated.Status)
	}
}

func TestUpdateTaskStatus_RejectsNonMember(t *testing.T) {
	db := setupTestDB(t)
	projectRepo := repository.NewProjectRepository(db)
	projectSvc := NewProjectService(projectRepo)
	taskRepo := repository.NewTaskRepository(db)
	taskSvc := NewTaskService(taskRepo)

	owner := createTestUser(t, db, "owner6@example.com")
	outsider := createTestUser(t, db, "outsider2@example.com")

	project, _ := projectSvc.CreateProject("Project W", "", owner.ID)
	task, _ := taskSvc.CreateTask(project.ID, owner.ID, "Private task", "", nil)

	_, err := taskSvc.UpdateTaskStatus(task.ID, outsider.ID, models.StatusDone)
	if err != ErrNotAuthorized {
		t.Errorf("expected ErrNotAuthorized, got: %v", err)
	}
}
