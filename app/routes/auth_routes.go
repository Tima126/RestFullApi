package routes

import (
	"app/handlers/auth"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router) {
	r.Post("/login", auth.RegisterHandler)
}
