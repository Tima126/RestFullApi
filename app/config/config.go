package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	DB_DSN     string
	JWT_Secret string
	AppPort    string
}

func LoadConfig() *Config {

	_ = godotenv.Load()

	cfg := &Config{
		AppEnv:     getEnv("APP_ENV", "dev"),
		DB_DSN:     getEnv("DB_DSN", "postgres://admin:12345@localhost:5434/restapi_db?sslmode=disable"),
		JWT_Secret: getEnv("JWT_SECRET", "supersecret"),
		AppPort:    getEnv("APP_PORT", "8080"),
	}

	return cfg
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
