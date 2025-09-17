package routes

import (
	jwt "app/jwt_token"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(jwtService *jwt.JWTService) *chi.Mux {
	r := chi.NewRouter()

	AuthRoutes(r, jwtService)
	UserRoutes(r, jwtService)

	return r
}
