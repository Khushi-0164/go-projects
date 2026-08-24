package main

import (
	"log/slog"
	"os"

	"invoice-saas/config"
	"invoice-saas/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found, relying on system environment variables")
	}

	db := config.ConnectDB()

	router := routes.SetupRouter(db)

	port := config.GetEnv("PORT", "8080")
	slog.Info("server starting", "port", port)
	if err := router.Run(":" + port); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
