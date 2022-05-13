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

var al aggregatedLogger

func init() {
	flags := log.LstdFlags | log.Lshortfile

	infoLogger := log.New(os.Stdout, "[INFO] ", flags)
	warnLogger := log.New(os.Stdout, "[WARN] ", flags)
	errorLogger := log.New(os.Stdout, "[ERR] ", flags)

	al = aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}
}

func Info(v ...interface{}) {
	al.infoLogger.Println(v...)
}

func Warn(v ...interface{}) {
	al.warnLogger.Println(v...)
}

func Error(v ...interface{}) {
	al.errorLogger.Println(v...)
}
