package logger

import (
	"log"
	"os"
)

type aggregatedLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

var logger aggregatedLogger

func init() {
	flags := log.LstdFlags | log.Lshortfile

	infoLogger := log.New(os.Stdout, "[INFO] ", flags)
	warnLogger := log.New(os.Stdout, "[WARN] ", flags)
	errorLogger := log.New(os.Stdout, "[ERR] ", flags)

	logger = aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}
}

func (l aggregatedLogger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l aggregatedLogger) Warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l aggregatedLogger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

func GetLogger() aggregatedLogger {
	return logger
}
