package cfg

import (
	"fmt"

	"github.com/caarlos0/env"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Config struct {
	UniStatsURL string `env:"UNI_STATS_URL,required"`

	Token string `env:"TOKEN,required"`
}

// NewConfig returns a new Config struct
func NewConfig() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		panic(fmt.Sprintf("Cant parse env vars: %+v", err))
	}

	if err := ValidateConfig(cfg); err != nil {
		panic(fmt.Sprintf("Invalid config: %+v", err))
	}

	return &cfg
}

func ValidateConfig(c Config) error {
	return validation.ValidateStruct(&c, validation.Field(&c.UniStatsURL, validation.Required, is.URL))
}
