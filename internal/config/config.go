package config

import (
	"fmt"
	"os"
)

type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SERVER_PORT string
	JWT_SECRET  string
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.DB_USER, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME,
	)
}

func Load() (*Config, error) {
	cfg := &Config{
		DB_HOST:     getEnv("DB_HOST", "localhost"),
		DB_PORT:     getEnv("DB_PORT", "5432"),
		DB_USER:     getEnv("DB_USER", "chat_user"),
		DB_PASSWORD: getEnv("DB_PASSWORD", "chat_password"),
		DB_NAME:     getEnv("DB_NAME", "chat_db"),
		SERVER_PORT: getEnv("SERVER_PORT", "8080"),
		JWT_SECRET:  getEnv("JWT_SECRET", ""),
	}

	return cfg, nil
}
