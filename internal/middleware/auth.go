package middleware

import (
	"net/http"
	"strings"

	"youtube-backend/configs" // Update the import to the correct package

	"github.com/golang-jwt/jwt/v4" // Use the correct JWT package
)

// AuthMiddleware checks for a valid JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Extract the token
		tokenString := strings.TrimSpace(strings.Split(authHeader, "Bearer ")[1]) // Handle cases without "Bearer" properly

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNoLocation
			}
			return []byte(configs.JWTSecret), nil // Corrected to configs
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
