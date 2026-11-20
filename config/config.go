package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config stores project settings
type Config struct {
	AppName string
	AppPort string
	AppEnv  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

// Load reads project settings
func Load() (*Config, error) {

	// Read .env file
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	// Create config
	config := &Config{
		AppName: os.Getenv("APP_NAME"),
		AppPort: os.Getenv("APP_PORT"),
		AppEnv:  os.Getenv("APP_ENV"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
	}

	return config, nil
}
