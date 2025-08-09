package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

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
		code := "internal_server_error"
		response.AbortWithStatusJSON(c, http.StatusInternalServerError, code, "大変申し訳ございません。システムに不具合が生じております。\n管理者にお問い合わせください。", nil)
	})
}
