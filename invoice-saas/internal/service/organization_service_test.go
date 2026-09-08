package service

import (
	"testing"

	"invoice-saas/internal/models"
)

func TestCreateOrganization_MakesCreatorOwner(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := NewOrganizationService(repo)

	org, err := svc.CreateOrganization("Acme Co", 1)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if org.ID == 0 {
		t.Errorf("expected org to have a non-zero ID after creation")
	}

	role, isMember := repo.FindMemberRole(org.ID, 1)
	if !isMember {
		t.Fatalf("expected creator to be a member of the org, but they are not")
	}
	if role != models.OrgRoleOwner {
		t.Errorf("expected role %q, got %q", models.OrgRoleOwner, role)
	}
}

func TestAddMember_RejectsNonOwnerNonAdmin(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := NewOrganizationService(repo)

	// Owner creates the org.
	org, err := svc.CreateOrganization("Acme Co", 1)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	// Manually add user 2 as a regular member (not owner/admin).
	if err := repo.AddMember(&models.OrgMember{
		OrganizationID: org.ID,
		UserID:         2,
		Role:           models.OrgRoleMember,
	}); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	// Register a target user to invite.
	repo.addUser(&models.User{ID: 3, Email: "target@example.com"})

	// User 2 (a regular member) tries to add someone — should be rejected.
	_, err = svc.AddMember(org.ID, 2, "target@example.com", models.OrgRoleMember)
	if err == nil {
		t.Fatalf("expected an error when a regular member tries to add someone, got nil")
	}
	if err != ErrNotAuthorized {
		t.Errorf("expected ErrNotAuthorized, got: %v", err)
	}
}

func TestAddMember_SucceedsForOwner(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := NewOrganizationService(repo)

	org, err := svc.CreateOrganization("Acme Co", 1)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	repo.addUser(&models.User{ID: 2, Email: "target@example.com"})

	member, err := svc.AddMember(org.ID, 1, "target@example.com", models.OrgRoleAdmin)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if member.Role != models.OrgRoleAdmin {
		t.Errorf("expected role %q, got %q", models.OrgRoleAdmin, member.Role)
	}
}

func TestListMyOrganizations_OnlyReturnsUsersOrgs(t *testing.T) {
	repo := newFakeOrgRepo()
	svc := NewOrganizationService(repo)

	svc.CreateOrganization("User 1's Org", 1)
	svc.CreateOrganization("User 2's Org", 2)

	orgs, err := svc.ListMyOrganizations(1)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(orgs) != 1 {
		t.Fatalf("expected user 1 to see exactly 1 org, got %d", len(orgs))
	}
	if orgs[0].Name != "User 1's Org" {
		t.Errorf("expected %q, got %q", "User 1's Org", orgs[0].Name)
	}
}
