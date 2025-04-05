package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort    string
	MongoURI    string
	MongoDBName string
	JikanAPIURL string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	return &Config{
		HTTPPort:    getEnv("HTTP_PORT", "8081"),
		MongoURI:    getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDBName: getEnv("MONGO_DB_NAME", "inventory"),
		JikanAPIURL: getEnv("JIKAN_API_URL", "https://api.jikan.moe/v4"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
