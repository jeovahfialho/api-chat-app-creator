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
		return nil, err
	}

	config := &Config{
		Port:         os.Getenv("PORT"),
		ClaudeAPIKey: os.Getenv("CLAUDE_API_KEY"),
	}

	fmt.Printf("Loaded API Key: %s\n", config.ClaudeAPIKey) // Log para debug

	return config, nil
}
