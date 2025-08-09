package middleware

import (
	"time"

	"github.com/forroots/nomity-admin-api-v3/internal/shared/contextx"
	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		contextx.GetLogger(c.Request.Context()).Info("[REQUEST START]")

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		contextx.GetLogger(c.Request.Context()).Info("[REQUEST END]",
			"status", status,
			"latency", latency,
		)
	}
}
