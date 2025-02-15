package http_handler

// import (
// 	"net/http"

// 	"github.com/barlus-engineer/barlus-api/Internal/dto"
// 	"github.com/barlus-engineer/barlus-api/pkg/text"
// 	"github.com/gin-gonic/gin"
// )

// var (

// )

// func ErrHandlerAuth(c *gin.Context, err text.AppError) {
// 	switch err.SvcErr {	
// 		default:
// 			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{
// 				Error: err.AppErr.Error(),
// 			})
// 	}
// }