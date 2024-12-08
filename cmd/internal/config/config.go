package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	ServerPort    string

	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
}

// GetEnv retrieves environment variables or returns a default value if not set.
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Init initializes the configuration by reading from environment variables.
func Init() *Config {
	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return &Config{
		ServerAddress: GetEnv("SERVER_ADDRESS", "http://localhost"),
		ServerPort:    GetEnv("SERVER_PORT", "8000"),
		DBHost:        GetEnv("DB_HOST", "localhost"),
		DBPort:        GetEnv("DB_PORT", "3306"),
		DBName:        GetEnv("DB_NAME", "go_rest_api"),
		DBUser:        GetEnv("DB_USER", "root"),
		DBPassword:    GetEnv("DB_PASSWORD", ""),
	}
}

// Global configuration instance
var Env = Init()
