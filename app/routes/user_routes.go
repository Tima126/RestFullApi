package routes

import (
	"app/handlers"
	"app/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTMiddleware)
		r.Get("/profil", handlers.ProfileHandler)
	})

}
