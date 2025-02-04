package config

import (
	"github.com/barlus-engineer/barlus-api/pkg/getenv"
	"github.com/barlus-engineer/barlus-api/pkg/logger"
	"github.com/barlus-engineer/barlus-api/pkg/text"
)

var config ConfigStrc

type ConfigStrc struct {
	Name    string `envkey:"SERVER_NAME" envdef:"Barlus API"`
	Version string `envkey:"SERVER_VERSION" envdef:"1.0 dev"`
	Release bool   `envkey:"SERVER_RELEASE" envdef:"false"`
	HTTP  struct {
		Host    string `envkey:"HTTP_HOST" envdef:"localhost"`
		Port    int    `envkey:"HTTP_PORT" envdef:"3250"`
	}
}

func LoadConfig() error {
	var (
		cfg ConfigStrc
	)

	if err := getenv.GetStruct(&cfg); err != nil {
		logger.Fatalf(text.ErrConfigGetenv.Error(), err)
	}

	config = cfg

	return nil
}

func GetConfig() *ConfigStrc {
	return &config
}