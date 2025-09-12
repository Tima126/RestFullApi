package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log - глобальный логгер
var Log = logrus.New()

// Init - инициализация логгера
func Init() {
	// установка уровня логирования и форматирования
	env := os.Getenv("APP_ENV")
	Log.SetLevel(logrus.DebugLevel)

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})

	// установка уровня логирования в зависимости от окружения
	switch env {
	case "dev":
		{
			Log.SetLevel(logrus.DebugLevel)
		}
	case "prod":
		{
			Log.SetLevel(logrus.InfoLevel)
		}
	default:
		{
			Log.SetLevel(logrus.WarnLevel)
		}
	}

	Log.SetOutput(os.Stdout)

}
