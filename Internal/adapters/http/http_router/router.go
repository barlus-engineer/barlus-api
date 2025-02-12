package http_router

import (
	"github.com/barlus-engineer/barlus-api/Internal/adapters/database"
	"github.com/barlus-engineer/barlus-api/Internal/adapters/http/http_handler"
	"github.com/barlus-engineer/barlus-api/Internal/adapters/repository"
	"github.com/barlus-engineer/barlus-api/Internal/core/services"
	"github.com/gin-gonic/gin"
)

func Route(api *gin.Engine) {
	api.NoRoute(http_handler.NoRoute)
	api.GET("/ping", http_handler.Ping)

	auth := api.Group("/auth")
	{
		db := database.GetDatabase()
		userRepo := &repository.UserRepo{}
		userRepo.AddDatabase(db)
		userService := services.NewUserService(userRepo)
		authHandler :=  http_handler.AuthHandler{}.NewAuthenHandler(*userService)

		auth.POST("/register", authHandler.Register)
	}
}