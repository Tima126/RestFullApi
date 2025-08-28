package main

import (
	"app/db"
	"app/logger"
	"app/routes"
	"net/http"
)

func main() {
	logger.Init()
	r := routes.RegisterRoutes()

	logger.Log.Info("Сервер запущен на :8080")
	db.Init()
	http.ListenAndServe(":8080", r)
}
