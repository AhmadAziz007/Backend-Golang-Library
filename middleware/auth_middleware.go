// middleware/auth_middleware.go
package middleware

import (
	"library-synapsis/helper"
	"library-synapsis/model/web"
	"net/http"
)

//import (
//	"library-synapsis/helper"
//	"library-synapsis/model/web"
//	"net/http"
//)

//func AuthMiddleware(next http.Handler, requiredRoles ...string) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		authHeader := r.Header.Get("Authorization")
//		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
//			http.Error(w, "Forbidden", http.StatusForbidden)
//			return
//		}
//
//		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
//		token, err := helper.ValidateJWT(tokenStr)
//		if err != nil || !token.Valid {
//			http.Error(w, "Unauthorized", http.StatusUnauthorized)
//			return
//		}
//
//		claims, _ := token.Claims.(jwt.MapClaims)
//		userRole := claims["role"].(string)
//
//		// Check if role is authorized
//		isAuthorized := false
//		for _, role := range requiredRoles {
//			if role == userRole {
//				isAuthorized = true
//				break
//			}
//		}
//		if !isAuthorized {
//			http.Error(w, "Forbidden", http.StatusForbidden)
//			return
//		}
//
//		// Add claims to context
//		ctx := context.WithValue(r.Context(), "claims", claims)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	})
//}

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
