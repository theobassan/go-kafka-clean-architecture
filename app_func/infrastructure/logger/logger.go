package logger

import (
	"log"
	"os"

	"go-kafka-clean-architecture/app_func/logger"

	"github.com/go-errors/errors"
)

func NewLogger() (*log.Logger, error) {

	mainPath, err := os.Getwd()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	logPath := mainPath + "/log/error.log"
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.Wrap(err, 1)
	}

	return log.New(f, "", log.LstdFlags|log.Lshortfile), nil
}

func NewDebugLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func Error(logger *log.Logger) logger.LoggerError {
	return func(i interface{}) {
		logger.Println(i)
	}
}
