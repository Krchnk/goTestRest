package psql

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Host   string `env:"HOST"`
	Port   string `env:"PORT"`
	DBName string `env:"DB_NAME"`
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	Path   string `env:"PATH"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Parse() error {
	if err := env.Parse(c); err != nil {
		return err
	}
	return nil
}
