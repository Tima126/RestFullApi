package routes

import "github.com/go-chi/chi/v5"

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	AuthRoutes(r)
	UserRoutes(r)

	return r
}
