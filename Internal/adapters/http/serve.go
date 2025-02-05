package http

import (
	"fmt"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/http/router"
	"github.com/barlus-engineer/barlus-api/config"
	"github.com/barlus-engineer/barlus-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RunHTTPServe() {
	var (
		cfg        = config.GetConfig()
		serverAddr = fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)
		serve      *gin.Engine
	)

	if cfg.Release {
		gin.SetMode(gin.ReleaseMode)
		serve = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		serve = gin.Default()
	}

	router.Route(serve)

	logger.Info("HTTP server is starting on: ", serverAddr)
	if err := serve.Run(serverAddr); err != nil {
		logger.Crash(err)
	}
}
