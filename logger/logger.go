package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/markbates/going/defaults"
)

var logger *logrus.Logger

// GetLogger is used to return a singleton object of the logrus logger
func GetLogger() *logrus.Logger {

	if logger == nil {
		logger = logrus.New()
		var ENV = defaults.String(os.Getenv("GO_EUN_ENV"), "development")
		logger.Out = os.Stdout

		customFormatter := new(logrus.JSONFormatter)
		customFormatter.TimestampFormat = "2006-01-02 15:04:05"
		logger.Formatter = customFormatter

		if ENV == "development" {
			logger.Level = logrus.DebugLevel
		} else {
			logger.Level = logrus.InfoLevel
		}
	}

	return logger
}