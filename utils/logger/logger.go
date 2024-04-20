package logger

import (
	"log"
	"os"
)

type MyLogger struct {
	Instance *log.Logger
}

func loggerInit() {
	logger = &MyLogger{
		Instance: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func getLogger() *MyLogger {
	if logger == nil {
		loggerInit()
	}
	return logger
}

func Println(v ...interface{}) {
	getLogger().Instance.Println(v...)
}

func Printf(format string, v ...interface{}) {
	getLogger().Instance.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	getLogger().Instance.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	getLogger().Instance.Fatalf(format, v...)
}

func Panic(v ...interface{}) {
	getLogger().Instance.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	getLogger().Instance.Panicf(format, v...)
}

var logger *MyLogger
