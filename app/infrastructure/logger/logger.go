package logger

import (
	"log"
	"os"

	"go-kafka-clean-architecture/app/logger"

	"github.com/go-errors/errors"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger() (logger.Logger, error) {

	mainPath, err := os.Getwd()
	if !errors.Is(err, nil) {
		return nil, errors.Wrap(err, 1)
	}

	logPath := mainPath + "/log/error.log"
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.Wrap(err, 1)
	}

	logger := log.New(f, "", log.LstdFlags|log.Lshortfile)
	return &Logger{logger}, nil
}

//init .
func NewDebugLogger() logger.Logger {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	return &Logger{logger}
}

func (logger *Logger) Error(i interface{}) {
	logger.logger.Println(i)
}
