package middleware

import (
	"errors"
	"net/http"

	"github.com/forroots/nomity-admin-api-v3/internal/application"
	"github.com/forroots/nomity-admin-api-v3/internal/interface/response"
	"github.com/forroots/nomity-admin-api-v3/internal/shared/contextx"
	"github.com/gin-gonic/gin"
)

func NewErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // ハンドラ実行

		errs := c.Errors
		if len(errs) == 0 {
			return
		}
		// 最初のエラーのみ対応（複数エラーを扱いたければfor-loop）
		err := errs[0].Err

		logger := contextx.GetLogger(c.Request.Context())

		var appErr *application.AppError
		if errors.As(err, &appErr) {
			// 独自AppError型に対応
			logger.Error("application error occurred",
				"code", appErr.Code,
				"message", appErr.Message,
				"status", appErr.HttpStatus,
				"error", appErr.Err,
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
			)

			code := appErr.Code
			response.AbortWithStatusJSON(c, appErr.HttpStatus, code, appErr.Message, nil)
		} else {
			// その他は500で返す
			if logger != nil {
				logger.Error("unexpected error occurred",
					"error", err,
					"method", c.Request.Method,
					"path", c.Request.URL.Path,
				)
			}

			code := application.APP_ERROR_INTERNAL_SERVER_ERROR
			message := "申し訳ございません。予期せぬエラーが発生しました。後ほど再度お試しください。"
			response.AbortWithStatusJSON(c, http.StatusInternalServerError, code, message, nil)
		}
	}
}
