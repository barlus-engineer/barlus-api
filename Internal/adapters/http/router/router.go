package router

import (
	"github.com/barlus-engineer/barlus-api/Internal/adapters/http/handler"
	"github.com/gin-gonic/gin"
)

func Route(api *gin.Engine) {
	api.NoRoute(handler.NoRoute)
	api.GET("/ping", handler.Ping)

	authen := api.Group("/authen")
	{
		var authen_handler handler.Authen
		authen.POST("/register", authen_handler.Register)
	}
}