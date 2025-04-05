package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                string
	InventoryServiceURL string
	OrderServiceURL     string
	JWTSecret           string
	LogLevel            string
}

func Load() (*Config, error) {
	// Load .env file if it exists
	_ = godotenv.Load()

	port := getEnv("PORT", "8080")

	return &Config{
		Port:                port,
		InventoryServiceURL: getEnv("INVENTORY_SERVICE_URL", "http://inventory-service:8081"),
		OrderServiceURL:     getEnv("ORDER_SERVICE_URL", "http://order-service:8082"),
		JWTSecret:           getEnv("JWT_SECRET", "default-secret-key"),
		LogLevel:            getEnv("LOG_LEVEL", "debug"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
