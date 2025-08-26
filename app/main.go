package main

import (
	"app/handlers"
	"app/logger"
	"app/middleware"
	"net/http"
)

func main() {

	logger.Init()

	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.Handle("/profile", middleware.JWTMiddleware((http.HandlerFunc(handlers.ProfileHandler))))

	logger.Log.Info("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", mux)

}
