package main

import (
	"catalog-api/config"
	"catalog-api/internal/models"
	"catalog-api/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	db := config.ConnectDB()
	cache := config.ConnectRedis()

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	router := routes.SetupRouter(db, cache)

	port := config.GetEnv("PORT", "8080")
	log.Printf("Server starting on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
