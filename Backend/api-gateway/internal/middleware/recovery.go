package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Recovery middleware handles panics and returns 500 errors
func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, "Internal Server Error: "+err)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
