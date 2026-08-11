package config

import "os"

func Getenv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
func JWTSecret() []byte {
	return []byte(Getenv("JWT_SECRET", "change-me-in-production"))
}
