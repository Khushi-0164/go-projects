package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"invoice-saas/config"
	"invoice-saas/internal/repository"
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
	cache := config.ConnectRedis()

	summaryRepo := repository.NewSummaryRepository(db)
	cachedSummaryRepo := repository.NewCachedSummaryRepository(summaryRepo, cache)

	pool := worker.NewPool(db, cachedSummaryRepo, 3, 100)
	router := routes.SetupRouter(db, cachedSummaryRepo, pool)
	port := config.GetEnv("PORT", "8080")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		slog.Info("server starting", "port", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server failed", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutdown signal received, starting graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("graceful shutdown failed", "error", err)
		os.Exit(1)
	}

	slog.Info("server shut down cleanly")
}
