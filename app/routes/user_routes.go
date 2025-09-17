package routes

import (
	"app/handlers"
	jwt "app/jwt_token"
	"app/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, jwtService *jwt.JWTService) {
	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return middleware.JWTMiddleware(jwtService, next)
		})
		r.Get("/profil", handlers.ProfileHandler)
	})
}
