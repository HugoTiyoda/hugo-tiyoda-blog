package middleware

import (
	"blog/application/ports"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(sessionService ports.SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		session, err := sessionService.ValidateSession(tokenString)
		if err != nil {
			switch err.Error() {
			case "session expired":
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "session expired"})
			case "session not found":
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			}
			return
		}

		c.Set("session", session)
		c.Set("authorId", session.AuthorId)

		c.Next()
	}
}
