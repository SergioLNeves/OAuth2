package config

import (
	"fmt"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

func NewConfig() (*Config, error) {
	var config Config

	if err := loadConfigs(&config); err != nil {
		return nil, fmt.Errorf("load environment variables: %w", err)
	}

	return &config, nil
}

func loadConfigs(config *Config) error {
	_ = godotenv.Load()

	if _, err := env.UnmarshalFromEnviron(config); err != nil {
		return fmt.Errorf("load environment variables: %w", err)
	}

	return nil
}
