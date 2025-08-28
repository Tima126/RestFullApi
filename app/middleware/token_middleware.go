package middleware

import (
	jwt "app/jwt_token"
	"context"
	"net/http"
)

type contextKey string

const userCtxKey = contextKey("user")
const roleCtxKey = contextKey("role")

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := jwt.ExtractTokenFromHeader(r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := jwt.ParseToken(tokenString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, claims.UserID)
		ctx = context.WithValue(r.Context(), roleCtxKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
