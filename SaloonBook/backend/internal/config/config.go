package config

import (
	"log"
	"os"
)

type Config struct {
	Port          string
	DatabaseURL   string
	Environment   string
	ServeFrontend bool
	FrontendDist  string
}

func Load() *Config {
	cfg := &Config{
		Port:          getEnv("PORT", "5000"),
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		Environment:   getEnv("NODE_ENV", "development"),
		ServeFrontend: getEnv("SERVE_FRONTEND", "0") == "1",
		FrontendDist:  getEnv("FRONTEND_DIST", "../frontend/web/dist"),
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL must be set")
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
