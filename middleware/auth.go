package middleware

import (
	"ghgist-blog/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const RoleKey contextKey = "role"

// authentication middleware //remember c.abort() stops the execution of further middleware/handlers
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Check Bearer format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract user ID from claims
		userID, ok := (*claims)["user_id"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Extract the role of the user
		role, ok := (*claims)["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Store user info in Gin context for handlers to use
		c.Set("user_id", userID)
		c.Set("role", role)

		// Continue to next middleware/handler
		c.Next()
	}
}

// the middleware that extracts the role of the user from the response
func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get role from context (set by AuthMiddleware)
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No role found in context"})
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok || roleStr != "publisher" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You cannot access this resource"})
			c.Abort()
			return
		}

		// Continue to next middleware/handler
		c.Next()
	}
}
