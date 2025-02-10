package config

import (
	"errors"

	"github.com/barlus-engineer/barlus-api/pkg/getenv"
	"github.com/barlus-engineer/barlus-api/pkg/logger"
)

var (
	config ConfigStrc

	ErrGetenv = errors.New("config: error getting env:\n%v")
)

type ConfigStrc struct {
	Name    string `envkey:"SERVER_NAME" envdef:"Barlus API"`
	Version string `envkey:"SERVER_VERSION" envdef:"v0.1.1"`
	Release bool   `envkey:"SERVER_RELEASE" envdef:"false"`
	HTTP    struct {
		Host string `envkey:"HTTP_HOST" envdef:"localhost"`
		Port int    `envkey:"HTTP_PORT" envdef:"3250"`
	}
	Cache struct {
		RedisURL string `envkey:"REDIS_URL"`
		CacheTime int `envkey:"CACHE_TIME" envdef:"1"`
	}
	Database struct {
		PostgresURL string `envkey:"POSTGRES_URL"`
	}
}

func LoadConfig() error {
	var (
		cfg ConfigStrc
	)

	if err := getenv.GetStruct(&cfg); err != nil {
		logger.Crashf(ErrGetenv.Error(), err)
	}

	config = cfg

	return nil
}

func GetConfig() *ConfigStrc {
	return &config
}
