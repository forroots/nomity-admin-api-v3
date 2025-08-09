package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		getLogger(c).Info("[REQEST START]")

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		getLogger(c).Info("[REQUEST END]",
			"status", status,
			"latency", latency,
		)
	}
}
