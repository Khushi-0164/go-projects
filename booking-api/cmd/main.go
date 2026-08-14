package main

import (
	"booking-api/config"
	"booking-api/internal/models"
	"booking-api/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	db := config.ConnectDB()

	if err := db.AutoMigrate(&models.User{}, &models.Resource{}, &models.Booking{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := routes.SetupRouter(db)
	port := config.GetEnv("PORT", "8080")
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
