package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Mongo_URI  string
	Mongo_DB   string
	JWT_Secret string
}

func LoadEnv() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to load env file: %w", err)
	}

	cfg := Config{
		Mongo_URI:  strings.TrimSpace(os.Getenv("MONGO_DB_URI")),
		Mongo_DB:   strings.TrimSpace(os.Getenv("MONGO_DB")),
		JWT_Secret: strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}

	// Validate required envs
	if cfg.Mongo_URI == "" {
		return Config{}, fmt.Errorf("MONGO_DB_URI is required")
	}

	if cfg.Mongo_DB == "" {
		return Config{}, fmt.Errorf("MONGO_DB is required")
	}

	if cfg.JWT_Secret == "" {
		return Config{}, fmt.Errorf("JWT_SECRET is required")
	}

	return cfg, nil
}
