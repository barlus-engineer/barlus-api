package main

import (
	"os"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/cache"
	"github.com/barlus-engineer/barlus-api/Internal/adapters/database"
	"github.com/barlus-engineer/barlus-api/Internal/adapters/http"
	"github.com/barlus-engineer/barlus-api/config"
	"github.com/barlus-engineer/barlus-api/pkg/logger"
)

func main() {
	logger.Info("Server is starting...")
	logger.Infof("Server PID: %d", os.Getpid())

	LoadResource()

	http.RunHTTPServe()
}

func LoadResource() {
	if err := config.LoadConfig(); err != nil {
		logger.Crashf("config: %v", err)
	}
	logger.Info("Config loaded successfully")

	if err := cache.RedisConnect(); err != nil {
		logger.Crashf("redis: %v", err)
	}
	logger.Info("Connected to Redis successfully")

	if err := database.PostgresConnect(); err != nil {
		logger.Crashf("postgres: %v", err)
	}
	logger.Info("Connected to Database successfully")
}
