package middleware

import (
	jwt "app/jwt_token"
	"context"
	"net/http"
)

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

		ctx := context.WithValue(r.Context(), "user", claims.Login)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
