package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

type Config struct {
	AppEnv             string `env:"APP_ENV" envDefault:"local"`
	ServerPort         string `env:"SERVER_PORT" envDefault:"8445"`
	GeoIP2DatabasePath string `env:"GEO_IP_2_DATABASE_PATH"`
	MemcacheHost       string `env:"MEMCACHE_HOST"`
}

func NewConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
		return Config{}, err
	}

	return cfg, nil
}
