package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSON(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, SuccessResponse{
		Message: msg,
		Data:    data,
	})
}

func ErrorToMiddleware(c *gin.Context, err error) {
	c.Error(err)
}

func AbortWithStatusJSON(c *gin.Context, status int, code, msg string, data any) {
	c.AbortWithStatusJSON(status, ErrorResponse{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}
