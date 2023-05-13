package cfg

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	// The port to listen on
	Port string `env:"PORT"`
	// The database connection string
	MongoDB string `env:"MONGODB"`
}

// NewConfig returns a new Config struct
func NewConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("Cant parse env vars: %+v", err))
	}

	if cfg.MongoDB == "" || cfg.Port == "" {
		panic("Missing required env vars")
	}

	return &cfg
}
