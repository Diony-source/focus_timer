package utils

import (
	"os"
	"github.com/sirupsen/logrus"
)

// Logger yapılandırması
type Logger struct {
	*logrus.Logger
}

// GetLogger initializes and returns a logger instance
func GetLogger() *Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	return &Logger{logger}
}
