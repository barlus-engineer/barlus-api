package http_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": "router not found",
	})
}