package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/forroots/nomity-admin-api-v3/internal/application"
	"github.com/forroots/nomity-admin-api-v3/internal/interface/response"
	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		stack := string(debug.Stack())

		msg := fmt.Sprintf("%v", recovered)

		logger := getLogger(c)

		logger.Error("panic recovered",
			"error", msg,
			"stacktrace", stack,
		)

		// JSONで返す
		code := application.APP_ERROR_INTERNAL_SERVER_ERROR
		response.AbortWithStatusJSON(c, http.StatusInternalServerError, code, "申し訳ございません。予期せぬエラーが発生しました。後ほど再度お試しください。", nil)
	})
}
