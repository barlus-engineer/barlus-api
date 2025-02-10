package handler

import (
	"net/http"

	"github.com/barlus-engineer/barlus-api/Internal/adapters/repository"
	"github.com/barlus-engineer/barlus-api/Internal/core/model"
	"github.com/barlus-engineer/barlus-api/Internal/core/services"
	"github.com/gin-gonic/gin"
)

var (
	TextErrShouldBindBodyWithJSON = "Failed to bind request body with JSON"
)

type Authen struct {
	svc  services.User
	form model.User
}

func (p Authen) Register(c *gin.Context) {
	var (
		err error
	)
	if err = c.ShouldBindBodyWithJSON(&p.form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": TextErrShouldBindBodyWithJSON,
		})
		return
	}
	p.svc.AddData(p.form)
	if err = p.svc.Register(); err != nil {
		errorHandler(c, err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "register success",
	})
}

func errorHandler(c *gin.Context, err error) {
	switch err {
		case repository.ErrUnableCreateUser:
			c.JSON(http.StatusInternalServerError, gin.H{"error": repository.ErrUnableCreateUser.Error()})
		case repository.ErrUserExists:
			c.JSON(http.StatusBadRequest, gin.H{"error": repository.ErrUserExists.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}