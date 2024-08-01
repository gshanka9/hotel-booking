package utility

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
)

func InitLogger() *logrus.Logger {
	logger := logrus.New()
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetReportCaller(true)
	return logger
}

func LogWithStack(logger *logrus.Logger, msg interface{}) error {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	logger.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Error(msg)
	if err, ok := msg.(error); ok {
		return err
	}
	return fmt.Errorf("%v", msg)
}
