package middleware

import (
	"log/slog"

	"github.com/forroots/nomity-admin-api-v3/internal/interface/common"
	"github.com/forroots/nomity-admin-api-v3/internal/shared/contextx"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCORSHandler(allowOrigins []string, allowHeaders []string, logging bool) gin.HandlerFunc {
	if logging {
		return NewCORSWithLoggingHandler(allowOrigins, allowHeaders)
	}
	return newCORSHandler(allowOrigins, allowHeaders)
}

func newCORSHandler(allowOrigins []string, allowHeaders []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = allowOrigins

	allowHeaders = append(allowHeaders, common.XRequestIDHeader)
	config.AllowHeaders = append(config.AllowHeaders, allowHeaders...)
	return cors.New(config)
}

func NewCORSWithLoggingHandler(allowOrigins []string, allowHeaders []string) gin.HandlerFunc {
	corsHandler := newCORSHandler(allowOrigins, allowHeaders)

	return func(c *gin.Context) {
		logger := contextx.GetLogger(c.Request.Context())

		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			logger.Info("CORS request", slog.String("origin", origin), slog.String("method", c.Request.Method))
		}
		corsHandler(c)
	}
}
