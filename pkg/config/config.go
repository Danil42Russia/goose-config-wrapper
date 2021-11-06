package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Goose GooseConfig
}

type GooseConfig struct {
	Driver   string
	DBString string
	Options  GooseOptions
}

type GooseOptions struct {
	Dir          string
	Table        string
	Verbose      bool
	CertFile     string
	Sequential   bool
	AllowMissing bool
}

func Init() (*Config, error) {
	viper.SetConfigName(".goose")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error when reading config: %s", err)
	}

	cfg := new(Config)
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, fmt.Errorf("error when unmarshal config: %s", err)
	}

	return cfg, nil
}
