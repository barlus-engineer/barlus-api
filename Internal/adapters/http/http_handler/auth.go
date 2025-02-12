package http_handler

import (
	"net/http"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/repository"
	"github.com/barlus-engineer/barlus-api/Internal/core/services"
	"github.com/barlus-engineer/barlus-api/Internal/dto"
	"github.com/gin-gonic/gin"
)

var (
	TextErrShouldBindBodyWithJSON = "Failed to bind request body with JSON"
)

type AuthHandler struct {
	svc services.UserService
	regForm dto.UserRegisterRequest
}

func (p AuthHandler) NewAuthenHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{
		svc: userService,
	}
}

func (p AuthHandler) Register(c *gin.Context) {
	var (
		err error
	)
	if err = c.ShouldBindBodyWithJSON(&p.regForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": TextErrShouldBindBodyWithJSON,
		})
		return
	}
	if err = p.svc.Register(p.regForm); err != nil {
		registerErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "register success",
	})
}

func registerErrorHandler(c *gin.Context, err error) {
	switch err {
		case repository.ErrUnableCreateUser:
			c.JSON(http.StatusInternalServerError, gin.H{"error": repository.ErrUnableCreateUser.Error()})
		case repository.ErrUserExists:
			c.JSON(http.StatusBadRequest, gin.H{"error": repository.ErrUserExists.Error()})
		case repository.ErrUserEmailExists:
			c.JSON(http.StatusBadRequest, gin.H{"error": repository.ErrUserEmailExists.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}