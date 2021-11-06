package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

func Init() (*Config, error) {
	viper.SetConfigName(".goose")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error when reading config: %w", err)
	}

	cfg := NewDefault()
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("error when unmarshal config: %w", err)
	}

	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("error when validation config: %w", err)
	}

	return cfg, nil
}

func validateConfig(cfg *Config) error {
	g := cfg.Goose
	if g.Driver == "" {
		return errors.New("the driver name cannot be empty")
	}

	if g.DBString == "" {
		return errors.New("the dbstring cannot be empty")
	}

	return nil
}
