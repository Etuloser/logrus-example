package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func New() *logrus.Logger {
	var logger = logrus.New()

	logger.Out = os.Stdout

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	return logger
}
