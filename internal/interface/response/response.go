package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResponse(message string, data any) SuccessResponse {
	return SuccessResponse{
		Message: message,
		Data:    data,
	}
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func NewErrorResponse(code, message string, data any) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func JSON(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, NewSuccessResponse(msg, data))
}

func ErrorToMiddleware(c *gin.Context, err error) {
	c.Error(err)
}

func AbortWithStatusJSON(c *gin.Context, status int, code, msg string, data any) {
	c.AbortWithStatusJSON(status, NewErrorResponse(code, msg, data))
}
