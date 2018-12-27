package config

import "github.com/caarlos0/env"

type Config struct {
	Port          int    `env:"PORT" envDefault:"3000"`
	SessionSecret string `env:"SESSION_SECRET"`
	DBAddr        string `env:"DB_ADDR"`
	UIDomain      string `env:"UI_DOMAIN"`
}

func New() *Config {
	conf := &Config{}
	err := env.Parse(conf)
	if err != nil {
		panic(err)
	}
	return conf
}
