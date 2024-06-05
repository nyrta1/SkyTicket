package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP HTTPConfig
		PG   PGConfig
	}

	HTTPConfig struct {
		Port string `env:"HTTP_PORT"`
	}

	PGConfig struct {
		URL string `env:"PG_URL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig("./.env", cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("failed to load environment variables: %w", err)
	}

	return cfg, nil
}
