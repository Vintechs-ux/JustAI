package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DiscordToken string
	GroqAPIKey   string
	SentinelRole string
	PhantomRole  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env tidak ditemukan, gunakan environment variabel system")
	}

	cfg := &Config{
		DiscordToken: os.Getenv("DISCORD_BOT_TOKEN"),
		GroqAPIKey:   os.Getenv("GROQ_API_KEY"),
		SentinelRole: os.Getenv("SENTINEL_ROLE_NAME"),
		PhantomRole:  os.Getenv("PHANTOM_ROLE_NAME"),
	}

	if cfg.DiscordToken == "" {
		log.Fatal("DISCORD_BOT_TOKEN tidak ditemukan")
	}

	if cfg.GroqAPIKey == "" {
		log.Fatal("GROQ_API_KEY tidak ditemukan")
	}

	return cfg
}
