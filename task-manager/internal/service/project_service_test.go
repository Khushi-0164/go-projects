package service

import (
	"testing"

	"task-manager/internal/models"
	"task-manager/internal/repository"
)

func TestCreateProject_MakesCreatorOwner(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewProjectRepository(db)
	svc := NewProjectService(repo)

	user := createTestUser(t, db, "owner@example.com")

	project, err := svc.CreateProject("Test Project", "a description", user.ID)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if project.ID == 0 {
		t.Errorf("expected project to have a non-zero ID after creation")
	}

	if project.OwnerID != user.ID {
		t.Errorf("expected OwnerID %d, got %d", user.ID, project.OwnerID)
	}

	// Verify the ProjectMember row was actually created with role "owner"
	role, isMember := repo.FindMemberRole(project.ID, user.ID)
	if !isMember {
		t.Fatalf("expected creator to be a member of the project, but they are not")
	}
	if role != models.ProjectRoleOwner {
		t.Errorf("expected role %q, got %q", models.ProjectRoleOwner, role)
	}
}

func TestAddMember_RejectsNonOwnerNonAdmin(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewProjectRepository(db)
	svc := NewProjectService(repo)

	owner := createTestUser(t, db, "owner2@example.com")
	regularMember := createTestUser(t, db, "member@example.com")
	target := createTestUser(t, db, "target@example.com")

	project, err := svc.CreateProject("Another Project", "", owner.ID)
	if err != nil {
		t.Fatalf("setup failed: could not create project: %v", err)
	}

	// Manually add regularMember with role "member" (not owner/admin)
	if err := repo.AddMember(&models.ProjectMember{
		ProjectID: project.ID,
		UserID:    regularMember.ID,
		Role:      models.ProjectRoleMember,
	}); err != nil {
		t.Fatalf("setup failed: could not add regular member: %v", err)
	}

	// regularMember tries to add target — should be rejected
	_, err = svc.AddMember(project.ID, regularMember.ID, target.Email, models.ProjectRoleMember)
	if err == nil {
		t.Fatalf("expected an error when a regular member tries to add someone, got nil")
	}
	if err != ErrNotAuthorized {
		t.Errorf("expected ErrNotAuthorized, got: %v", err)
	}
}
