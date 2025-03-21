package logger

import (
	"log"
	"os"
)

type aggregatedLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
	fatalLogger *log.Logger
}

var al aggregatedLogger

func init() {
	flags := log.LstdFlags | log.LUTC

	infoLogger := log.New(os.Stdout, "[INFO] ", flags)
	warnLogger := log.New(os.Stdout, "[WARN] ", flags)
	errorLogger := log.New(os.Stdout, "[ERROR] ", flags)
	fatalLogger := log.New(os.Stdout, "[FATAL] ", flags)

	al = aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
		fatalLogger: fatalLogger,
	}
}

func Info(v ...interface{}) {
	al.infoLogger.Println(v...)
}

func Infof(format string, v ...interface{}) {
	al.infoLogger.Printf(format+"\n", v...)
}

func Warn(v ...interface{}) {
	al.warnLogger.Println(v...)
}

func Warnf(format string, v ...interface{}) {
	al.warnLogger.Printf(format+"\n", v...)
}

func Error(v ...interface{}) {
	al.errorLogger.Println(v...)
}

func Errorf(format string, v ...interface{}) {
	al.errorLogger.Printf(format+"\n", v...)
}

func Fatal(v ...interface{}) {
	al.fatalLogger.Fatalln(v...)
}

func Fatalf(format string, v ...interface{}) {
	al.fatalLogger.Fatalf(format+"\n", v...)
}
