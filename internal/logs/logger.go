package logs

import (
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var Logger *logrus.Logger

// NewLogger creates a new Logrus Logger and returns a pointer to logrus.Logger
func NewLogger() *logrus.Logger {

	var level logrus.Level
	level = LogLevel("info")
	logger := &logrus.Logger{
		Out:   os.Stdout,
		Level: level,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   true,
			TimestampFormat: "2009-06-03 11:04:075",
		},
	}
	Logger = logger
	return Logger
}

func LogLevel(lvl string) logrus.Level {
	switch lvl {
	case "info":
		return logrus.InfoLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	default:
		panic("Not supported")
	}
}
