// middleware/auth_middleware.go
package middleware

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"library-synapsis/helper"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler, requiredRoles ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := helper.ValidateJWT(tokenStr)
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)

		// Check if role is authorized
		isAuthorized := false
		for _, role := range requiredRoles {
			if role == userRole {
				isAuthorized = true
				break
			}
		}
		if !isAuthorized {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Add claims to context
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
