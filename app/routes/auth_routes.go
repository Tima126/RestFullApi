package routes

import (
	"app/handlers/auth"
	jwt "app/jwt_token"

	"github.com/go-chi/chi/v5"
)

// регистрация маршрутов аутентификации
func AuthRoutes(r chi.Router, jwtService *jwt.JWTService) {
	r.Post("/api/register", auth.RegisterHandler(jwtService))
	r.Post("/api/login", auth.LoginHandler(jwtService))
}
