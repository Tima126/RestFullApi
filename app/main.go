package main

import (
	"app/config"
	"app/db"
	jwt "app/jwt_token"
	"app/logger"
	"app/routes"
	"net/http"
)

func main() {
	// загрузка конфигурации
	cfg := config.LoadConfig()

	// инициализация логгера и базы данных
	logger.Init(cfg)
	db.Init(cfg)

	// инициализация JWT сервиса
	jwtService := jwt.NewJWTService(cfg)

	// регистрация роутов
	r := routes.RegisterRoutes(jwtService)

	logger.Log.Infof("Сервер запущен на :%s", cfg.AppPort)
	logger.Log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))

}
