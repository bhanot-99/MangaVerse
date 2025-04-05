package middleware

import (
	"log"
	"time"

	"api-gateway/internal/config"

	"github.com/gin-gonic/gin"
)

func Logger(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		log.Printf("[%s] %s %s %d %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Writer.Status(),
			latency,
		)
	}
}
