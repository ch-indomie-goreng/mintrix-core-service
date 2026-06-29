package config

import (
	"os"
	"time"
)

// Config holds all application configuration loaded from environment variables.
type Config struct {
	Port          string
	DBUrl         string
	JWTSecret     string
	JWTExpiry     time.Duration
	RefreshExpiry time.Duration
}

// Load reads configuration from environment variables with sensible defaults.
func Load() Config {
	return Config{
		Port:          getEnv("PORT", "8080"),
		DBUrl:         getEnv("DATABASE_URL", "postgresql://postgres:postgres@localhost:5432/mintrix?sslmode=disable"),
		JWTSecret:     getEnv("JWT_SECRET", "change-me-in-production"),
		JWTExpiry:     getEnvDuration("JWT_EXPIRY", 15*time.Minute),
		RefreshExpiry: getEnvDuration("REFRESH_EXPIRY", 7*24*time.Hour),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
	}
	return fallback
}
