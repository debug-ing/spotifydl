package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ClientID     string
	ClientSecret string
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		panic("No .env file found")
	}
	return &AppConfig{
		ClientID:     getEnvString("CLIENT_ID", ""),
		ClientSecret: getEnvString("CLIENT_SECRET", ""),
	}
}

func getEnvString(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
