package config

import "github.com/caarlos0/env"

type Config struct {
	PORT int `env:"PORT" envDefault:"8181"`
}

func New() (*Config, error) {
	config := &Config{}

	return config, env.Parse(config)
}
