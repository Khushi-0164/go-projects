package config

import (
	"context"
	"fmt"
	"log"
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
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "5432")
	user := GetEnv("DB_USER", "postgers")
	password := GetEnv("DB_PASSWORD", "postgers")
	name := GetEnv("DB_NAME", "catalogdb")
	sslmode := GetEnv("DB_SSLMODE", "disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, name, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Database connected successfully")
	return db
}

func ConnectRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: GetEnv("REDIS_ADDR", "localhost:6379"),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}
	log.Println("Redis connected successfully")
	return client
}
