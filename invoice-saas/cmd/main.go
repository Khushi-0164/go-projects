package main

import (
	"log/slog"
	"os"

	"invoice-saas/config"
	"invoice-saas/internal/routes"
	"invoice-saas/internal/worker"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Warn("no .env file found, relying on system environment variables")
	}

	stripe.Key = config.StripeSecretKey()

	db := config.ConnectDB()

	pool := worker.NewPool(db, 3, 100)
	router := routes.SetupRouter(db, pool)

	port := config.GetEnv("PORT", "8080")
	slog.Info("server starting", "port", port)
	if err := router.Run(":" + port); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
