package middleware

import (
	"strings"

	"mintrix-backend/internal/auth"
	"mintrix-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates the Bearer JWT and injects the userID into the context.
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer ") {
			utils.Unauthorized(c, "missing or malformed authorization header")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")
		claims, err := auth.ValidateAccessToken(tokenStr, jwtSecret)
		if err != nil {
			utils.Unauthorized(c, "invalid or expired token")
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
