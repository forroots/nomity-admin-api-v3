package middleware

import (
	"log/slog"

	"github.com/forroots/nomity-admin-api-v3/internal/interface/constants"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCORSHandler(allowOrigins []string, allowHeaders []string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = allowOrigins

	allowHeaders = append(allowHeaders, constants.XRequestIDHeader)
	config.AllowHeaders = append(config.AllowHeaders, allowHeaders...)
	return cors.New(config)
}

func NewCORSWithLoggingHandler(allowOrigins []string, allowHeaders []string) gin.HandlerFunc {
	corsHandler := NewCORSHandler(allowOrigins, allowHeaders)

	return func(c *gin.Context) {
		logger := getLogger(c)

		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			logger.Info("CORS request", slog.String("origin", origin), slog.String("method", c.Request.Method))
		}
		corsHandler(c)
	}
}
