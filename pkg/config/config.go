package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	ClaudeAPIKey string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found, using system environment variables\n")
	}

	config := &Config{
		Port:         os.Getenv("PORT"),
		ClaudeAPIKey: os.Getenv("CLAUDE_API_KEY"),
	}

	if config.ClaudeAPIKey == "" {
		return nil, fmt.Errorf("CLAUDE_API_KEY is not set")
	}

	fmt.Printf("Loaded API Key: %s\n", config.ClaudeAPIKey[:5]+"...")

	return config, nil
}
