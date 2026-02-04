package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken string
	GeminiKey    string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DiscordToken: os.Getenv("DISCORD_TOKEN"),
		GeminiKey:    os.Getenv("GEMINI_API_KEY"),
	}
}
