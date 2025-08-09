package middleware

import (
	"log/slog"

	"github.com/forroots/nomity-admin-api-v3/internal/interface/common"
	"github.com/forroots/nomity-admin-api-v3/internal/shared/contextx"
	"github.com/forroots/nomity-admin-api-v3/internal/utils"
	"github.com/gin-gonic/gin"
)

func EnrichContext(baseLogger *slog.Logger, sessionIdCookieName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストIDを生成
		reqID := utils.NewULIDString()
		// リクエストIDをヘッダーに設定
		c.Request.Header.Set(common.XRequestIDHeader, reqID)

		// セッションIDを取得
		sessionID, _ := c.Cookie(sessionIdCookieName)
		// ipアドレスを取得
		ip := c.ClientIP()
		// PathとMethodを取得
		method := c.Request.Method
		path := c.Request.URL.Path

		// ロガーに情報を追加
		enrichedLogger := baseLogger.With(
			"request_id", reqID,
			"session_id", sessionID,
			"ip_address", ip,
			"method", method,
			"path", path,
		)

		// ContextInfo構築
		info := &contextx.ContextInfo{
			Logger:    enrichedLogger,
			RequestID: reqID,
			SessionID: sessionID,
			IPAddress: ip,
			Method:    method,
			Path:      path,
		}

		// コンテキストに埋め込む
		ctx := contextx.WithContextInfo(c.Request.Context(), info)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
