package cfg

import (
	"fmt"

	"github.com/caarlos0/env"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Config struct {
	// The port to listen on
	Port string `env:"PORT"`
	// The database connection string
	MongoDB string `env:"MONGODB"`

	// EmailCheckService is the URL of the email check service
	EmailCheckService string `env:"EMAIL_CHECK_SERVICE"`
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
	return validation.ValidateStruct(&c, validation.Field(&c.Port, validation.Required),
		validation.Field(&c.MongoDB, validation.Required), validation.Field(&c.EmailCheckService, validation.Required, is.URL))
}
