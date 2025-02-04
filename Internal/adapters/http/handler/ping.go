package handler

import (
	"net/http"

	"github.com/barlus-engineer/barlus-api/config"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var (
		cfg = config.GetConfig()
	)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"server": cfg.Name,
		"version": cfg.Version,
	})
}