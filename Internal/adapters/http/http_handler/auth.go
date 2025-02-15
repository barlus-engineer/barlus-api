package http_handler

import (
	"net/http"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/repository"
	"github.com/barlus-engineer/barlus-api/Internal/core/services"
	"github.com/barlus-engineer/barlus-api/Internal/dto"
	"github.com/gin-gonic/gin"
)

var (
	TextErrShouldBindJSON = "Failed to bind request body with JSON"
)

type AuthHandler struct {
	svc services.UserService
}

func NewAuthHandler(userService services.UserService) *AuthHandler {
	return &AuthHandler{
		svc: userService,
	}
}

func (p *AuthHandler) Register(c *gin.Context) {
	var form dto.UserRegisterRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": TextErrShouldBindJSON,
		})
		return
	}
	if err := p.svc.Register(form); err != nil {
		ErrorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, dto.UserRegisterResponse{
		Message: "Register success",
	})
}

func (p *AuthHandler) UsernameAvail(c *gin.Context) {
	var form dto.UserUsernameAvailRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		ErrorHandler(c, err)
		return
	}
	if err := p.svc.UsernameAvail(form); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This username already exists",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "This username is available",
	})
}

func (p *AuthHandler) EmailAvail(c *gin.Context) {
	var form dto.UserEmailAvailRequest
	if err := c.ShouldBindJSON(&form); err != nil {
		ErrorHandler(c, err)
		return
	}
	if err := p.svc.EmailAvail(form); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This email already exists",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "This email is available",
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
