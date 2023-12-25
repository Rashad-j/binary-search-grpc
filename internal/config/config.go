package config

import "github.com/caarlos0/env"

type Config struct {
	Addr    string `env:"ADDR" envDefault:":8080"`
	MaxSize int    `env:"MAX_SIZE" envDefault:"1000000"`
}

func ReadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
