package main

import (
	"fmt"
	"os"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/handler"
	"github.com/barlus-engineer/barlus-api/config"
	"github.com/barlus-engineer/barlus-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Server is starting...")
	logger.Infof("Server PID: %d", os.Getpid())

	LoadResource()
	
	RunHTTPServe()
}

func LoadResource() {
	if err := config.LoadConfig(); err != nil {
		logger.Fatal(err)
	}
}

func RunHTTPServe() {
	var (
		cfg = config.GetConfig
		serveraddr = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
		router *gin.Engine
	)
	
	if cfg.Server.Release {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		router = gin.Default()
	}

	ping := router.Group("/ping")
	{
		ping.GET("/", handler.Ping)
	}

	logger.Info("HTTP server is starting on: ", serveraddr)
	router.Run(serveraddr)
}