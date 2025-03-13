package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	GitHubToken   string
	TelegramToken string
	DatabaseURL   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	config := &Config{
		GitHubToken:   os.Getenv("GITHUB_TOKEN"),
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
	}

	return config, nil
}
