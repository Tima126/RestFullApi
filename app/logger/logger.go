package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	env := os.Getenv("APP_ENV")
	Log.SetLevel(logrus.DebugLevel)

	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})

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
