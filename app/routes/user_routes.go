package routes

import (
	"app/handlers"
	jwt "app/jwt_token"
	"app/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, jwtService *jwt.JWTService) {
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTMiddleware(jwtService))
		r.Get("/profile", handlers.ProfileHandler)
	})
}
