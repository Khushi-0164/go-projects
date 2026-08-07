package service

import (
	"testing"

	"task-manager/config"
	"task-manager/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	// Tests run as a separate process from `go run ./cmd`, so .env
	// isn't loaded automatically — load it here too.
	_ = godotenv.Load("../../.env")

	db := config.ConnectDB()

	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM project_members")
	db.Exec("DELETE FROM projects")
	db.Exec("DELETE FROM users")

	return db
}
func createTestUser(t *testing.T, db *gorm.DB, email string) *models.User {
	t.Helper()

	user := &models.User{
		Email:        email,
		Name:         "Test User",
		PasswordHash: "irrelevant-for-these-tests",
	}
	if err := db.Create(user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}
	return user
}
