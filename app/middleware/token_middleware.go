package middleware

import (
	jwt "app/jwt_token"
	"context"
	"net/http"
)

type contextKey string

const userCtxKey = contextKey("user")
const roleCtxKey = contextKey("role")

// JWTMiddleware - middleware для проверки JWT токена
func JWTMiddleware(jwtService *jwt.JWTService, next http.Handler) http.Handler {

	// возврат анонимной функции-обработчика
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := jwt.ExtractTokenFromHeader(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// парсинг токена на клеймы
		claims, err := jwtService.ParseToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, claims.UserID)

		ctx = context.WithValue(r.Context(), roleCtxKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
