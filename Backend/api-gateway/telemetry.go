// api-gateway/telemetry.go
package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func telemetryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		log.Printf("%s %s %d %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}
