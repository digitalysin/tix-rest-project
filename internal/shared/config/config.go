package config

import (
	gconfig "github.com/digitalysin/goblog/config"
)

type (
	Configuration struct {
		HttpServerPort int `envconfig:"HTTP_SERVER_PORT" default:"3000"`

		// - database
		DatabaseConnectionString string `envconfig:"DATABASE_CONNECTION_STRING" required:"true"`

		// - cache
		CacheHost string `envconfig:"CACHE_HOST" default:"localhost:6379"`
	}
)

func New() (*Configuration, error) {
	var (
		cfg = Configuration{}
	)

	if err := gconfig.New(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
