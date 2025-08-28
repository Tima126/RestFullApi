package db

import (
	"app/logger"
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Init() {

	dsn := os.Getenv("DB_DSN")

	if dsn == "" {
		logger.Log.Fatal("DB_DSN environment variable is not set")
	}

	var err error
	Pool, err = pgxpool.New(context.Background(), dsn)

	if err != nil {
		logger.Log.Fatalf("Unable to create connection pool: %v", err)
	}

	logger.Log.Info("Connected to the database")
}
