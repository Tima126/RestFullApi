package main

import (
	"app/handlers"
	"app/middleware"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.Handle("/profile", middleware.JWTMiddleware((http.HandlerFunc(handlers.ProfileHandler))))

	fmt.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", mux)

}
