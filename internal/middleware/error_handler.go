package middleware

import (
	"log/slog"
	"net/http"

	"mintrix-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// ErrorHandler recovers from panics and returns a consistent JSON error response.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				slog.Error("panic recovered", "panic", r)
				utils.Error(c, http.StatusInternalServerError, "internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
