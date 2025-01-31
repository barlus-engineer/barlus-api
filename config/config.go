package config

import (
	"fmt"
	"strings"

	"github.com/barlus-engineer/barlus-api/pkg/logger"
	"github.com/barlus-engineer/barlus-api/pkg/text"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var GetConfig ConfigStrc

type ConfigStrc struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Server  struct {
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Release bool   `yaml:"release"`
	} `yaml:"server"`
}

func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		logger.Warningf("%s: %v", text.ErrDotenvLoad, err)
	}

	viper.SetConfigName("default.conf")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.BindEnv("server.host", "SERVER_HOST")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("server.release", "SERVER_RELEASE")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("%s: %v", text.ErrReadConfigFile, err)
	}

	var cfg ConfigStrc
	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("%s: %v", text.ErrDecodeConfig, err)
	}

	GetConfig = cfg

	return nil
}
