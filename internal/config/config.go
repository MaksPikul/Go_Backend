package config

import (
	"errors"
	"log/slog"
	"os"
)

type Config struct {
	Port  string
	DBUrl string
}

func LoadConfig() (*Config, error) {

	cfg := &Config{
		Port:  os.Getenv("PORT"),
		DBUrl: os.Getenv("DATABASE_URL_SQL"),
	}

	// prehaps a handle config errors function ?
	if cfg.Port == "" {
		slog.Warn("PORT not set in environment. Using default :8080")
		cfg.Port = "8080"
	}
	if cfg.DBUrl == "" {
		return nil, errors.New("DATABASE_URL_SQL is required")
	}

	return cfg, nil
}
