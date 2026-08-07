package main

import (
	"log"

	"task-manager/config"
	"task-manager/internal/models"
	"task-manager/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	db := config.ConnectDB()

	if err := db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.ProjectMember{},
		&models.Task{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := routes.SetupRouter(db)

	port := config.GetEnv("PORT", "8080")
	log.Printf("Server starting on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
