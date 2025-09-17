package logger

import (
	"app/config"
	"os"

	"github.com/sirupsen/logrus"
)

// Log - глобальный логгер
var Log = logrus.New()

// Init - инициализация логгера
func Init(cfg *config.Config) {
	// установка уровня логирования и форматирования
	env := cfg.AppEnv

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
