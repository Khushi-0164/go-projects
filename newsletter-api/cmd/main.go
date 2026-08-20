package main

import (
	"log/slog"
	"newsletter-api/config"
	"newsletter-api/internal/models"
	"newsletter-api/internal/routes"
	"newsletter-api/internal/worker"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found, relying on system environment variables")
	}

	db := config.ConnectDB()

	if err := db.AutoMigrate(&models.Subscriber{}); err != nil {
		slog.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}

	pool := worker.NewPool(3, 100)

	router := routes.SetupRouter(db, pool)

	port := config.GetEnv("PORT", "8080")
	slog.Info("server starting", "port", port)
	if err := router.Run(":" + port); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
