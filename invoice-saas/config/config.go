package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		GetEnv("DB_HOST", "localhost"),
		GetEnv("DB_PORT", "5432"),
		GetEnv("DB_USER", "postgres"),
		GetEnv("DB_PASSWORD", "postgres"),
		GetEnv("DB_NAME", "invoicesaas"),
		GetEnv("DB_SSLMODE", "disable"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}
	slog.Info("database connected successfully")
	return db
}

func JWTSecret() []byte {
	return []byte(GetEnv("JWT_SECRET", "change-me-in-production"))
}

func StripeSecretKey() string {
	return GetEnv("STRIPE_SECRET_KEY", "")
}

func StripeWebhookSecret() string {
	return GetEnv("STRIPE_WEBHOOK_SECRET", "")
}
func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_ADDR", "localhost:6379"),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Error("failed to connect to redis", "error", err)
		os.Exit(1)
	}
	slog.Info("redis connected successfully")
	return client
}
