package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("Failed to log to file, using default stderr: %v", err)
	}

	logger.SetOutput(file)

	return logger
}
