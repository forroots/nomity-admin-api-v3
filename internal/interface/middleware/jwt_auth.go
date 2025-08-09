package middleware

import (
	"net/http"
	"strings"

	"github.com/forroots/nomity-admin-api-v3/internal/application"
	"github.com/forroots/nomity-admin-api-v3/internal/interface/response"
	"github.com/forroots/nomity-admin-api-v3/internal/shared/contextx"
	"github.com/forroots/nomity-admin-api-v3/internal/utils/jwt"
	"github.com/gin-gonic/gin"
)

// const ContextKeyUserID = "user_id"

func NewJWTAuthMiddleware(jwtUtil *jwt.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := contextx.GetLogger(c.Request.Context())

		// AuthorizationヘッダーからBearerトークンを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			logger.Error("Authorization header is missing or invalid", "header", authHeader)
			response.AbortWithStatusJSON(
				c,
				http.StatusUnauthorized,
				application.APP_ERROR_UNAUTHORIZED,
				"[1]unauthorized request",
				nil,
			)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// JWTトークンをパースしてadminUserIdを取得
		adminUserId, err := jwtUtil.ParseToken(tokenStr)
		if err != nil {
			logger.Error("Failed to parse JWT token", "error", err)
			response.AbortWithStatusJSON(
				c,
				http.StatusUnauthorized,
				application.APP_ERROR_UNAUTHORIZED,
				"[2]unauthorized request",
				nil,
			)
			return
		}

		// コンテキストにadminUserIdをセット
		contextx.SetAdminUserID(c.Request.Context(), adminUserId)

		c.Next()
	}
}
