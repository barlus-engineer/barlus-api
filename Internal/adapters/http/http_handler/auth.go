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
}

func (p AuthHandler) NewAuthenHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{
		svc: userService,
	}
}

func (p AuthHandler) Register(c *gin.Context) {
	var (
		form dto.UserRegisterRequest
		err error
	)
	if err = c.ShouldBindBodyWithJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": TextErrShouldBindBodyWithJSON,
		})
		return
	}
	if err = p.svc.Register(form); err != nil {
		ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "register success",
	})
}

func (p AuthHandler) UsernameAvail(c *gin.Context) {
	var (
		form dto.UserUsernameAvailRequest
		err error
	)
	if err = c.ShouldBindBodyWithJSON(&form); err != nil {
		ErrorHandler(c, err)
		return
	}
	if err = p.svc.UsernameAvail(form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "This username is available",
		})
		return
	}
	c.JSON(http.StatusConflict, gin.H{
		"error": "This username is already exists",
	})
}

func ErrorHandler(c *gin.Context, err error) {
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