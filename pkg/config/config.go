package config

var defaultGooseOptions = GooseOptions{
	Dir:          ".",
	Table:        "goose_db_version",
	CertFile:     "",
	Verbose:      false,
	Sequential:   false,
	AllowMissing: false,
}

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
	CertFile     string
	Verbose      bool
	Sequential   bool
	AllowMissing bool `mapstructure:"allow-missing"`
}

func (c *Config) DefaultGooseOptions() GooseOptions {
	return defaultGooseOptions
}

func NewDefault() *Config {
	return &Config{
		Goose: GooseConfig{
			Options: defaultGooseOptions,
		},
	}
}
