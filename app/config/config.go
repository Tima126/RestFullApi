package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv     string
	DB_DSN     string
	JWT_Secret string
	APP_PORT   string
}

func LoadConfig() *Config {

	_ = godotenv.Load(".env.local")

	cfg := &Config{
		AppEnv:     getEnv("APP_ENV", ""),
		DB_DSN:     getEnv("DB_DSN", ""),
		JWT_Secret: getEnv("JWT_SECRET", ""),
		APP_PORT:   getEnv("APP_PORT", ""),
	}

	log.Printf("Конфигурация загружена: %+v\n", cfg)
	return cfg
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
