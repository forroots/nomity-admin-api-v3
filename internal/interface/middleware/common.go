package middleware

import (
	"log/slog"

	"github.com/forroots/nomity-admin-api-v3/internal/shared/contextx"
	"github.com/gin-gonic/gin"
)

func getLogger(c *gin.Context) *slog.Logger {
	// enriched logger があれば使う、なければdefaultのslogを使う
	if info, ok := contextx.GetContextInfo(c.Request.Context()); ok && info.Logger != nil {
		return info.Logger
	}
	return slog.Default()
}
