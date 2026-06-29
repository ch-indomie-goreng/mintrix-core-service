package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID injects a unique request identifier into every response header and context.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Header("X-Request-ID", id)
		c.Set("requestID", id)
		c.Next()
	}
}
