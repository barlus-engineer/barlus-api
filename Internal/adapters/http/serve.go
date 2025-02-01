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
		cfg = config.GetConfig()
		serverAddr = fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
		serve *gin.Engine
	)
	
	if cfg.Server.Release {
		gin.SetMode(gin.ReleaseMode)
		serve = gin.New()
	} else {
		gin.SetMode(gin.DebugMode)
		serve = gin.Default()
	}

	router.Route(serve)

	logger.Info("HTTP server is starting on: ", serverAddr)
	serve.Run(serverAddr)
}