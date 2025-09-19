package middleware

import (
	jwt "app/jwt_token"
	"context"
	"net/http"
)

type contextKey string

const UserCtxKey = contextKey("user")
const RoleCtxKey = contextKey("role")

func JWTMiddleware(jwtService *jwt.JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString, err := jwt.ExtractTokenFromHeader(r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			claims, err := jwtService.ParseToken(tokenString)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserCtxKey, claims.Login)
			ctx = context.WithValue(ctx, RoleCtxKey, claims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
