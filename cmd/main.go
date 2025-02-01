package main

import (
	"os"

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
		logger.Fatal(err)
	}
}