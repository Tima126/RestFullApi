package db

/**********************************************************
pool соединение с постгресс, строка подключение получаеться
из переменной окружение которое задаёться в docker-compose
***********************************************************/

import (
	"app/config"
	"app/logger"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init(cfg *config.Config) {

	dsn := cfg.DB_DSN

	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)

	if err != nil {
		logger.Log.Fatalf("Unable to create connection pool: %v", err)
	}

	logger.Log.Info("Connected to the database")
}
