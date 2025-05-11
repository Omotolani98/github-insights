package config

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/charmbracelet/log"
)

type APPENV struct {
	PORT string
	GITHUB_CLIENT_ID string
	GITHUB_CLIENT_SECRET string
}

func LoadEnv () APPENV {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("Could not load environment files")
	}

	appEnv := APPENV {
		PORT: getEnv("PORT", "8080"),
		GITHUB_CLIENT_ID: getEnv("GITHUB_CLIENT_ID", "nil"),
		GITHUB_CLIENT_SECRET: getEnv("GITHUB_CLIENT_SECRET", "nil"),
	}

	return appEnv
}

func getEnv (key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	log.Warnf("Environment Variable with key [%s] not set, Using Default: %s", key, fallback)
	return fallback
}
