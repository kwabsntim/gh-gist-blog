package middleware

import (
	"context"
	"ghgist-blog/utils"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const RoleKey contextKey = "role"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get token from Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check Bearer format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
			return
		}

		token := parts[1]

		// Validate token
		claims, err := utils.ValidateToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract user ID from claims
		userID, ok := (*claims)["user_id"].(string)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		//Extract the role of the user
		role, ok := (*claims)["role"].(string)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, RoleKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// the middleware that extracts the role of the user from the response
func RoleMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Context().Value(RoleKey).(string)
		if role != "publisher" {
			http.Error(w, "You cannot access this resource", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
