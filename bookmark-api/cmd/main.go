package main

import (
	"bookmark-api/config"
	"bookmark-api/internal/models"
	"bookmark-api/internal/routes"
	"log/slog"
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

	if err := db.AutoMigrate(&models.Bookmark{}); err != nil {
		slog.Error("failed to migrate database", "error", err)
		os.Exit(1)
	}

	router := routes.SetupRouter(db)

	port := config.Getenv("PORT", "8080")
	slog.Info("server starting", "port", port)
	if err := router.Run(":" + port); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
