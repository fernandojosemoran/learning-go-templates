package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

var (
	infoLogger  = log.New(os.Stdout, fmt.Sprintf("%s[INFO]%s  ", green, reset), log.Ldate|log.Ltime)
	warnLogger  = log.New(os.Stdout, fmt.Sprintf("%s[WARN]%s  ", yellow, reset), log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, fmt.Sprintf("%s[ERROR]%s ", red, reset), log.Ldate|log.Ltime)
	debugLogger = log.New(os.Stdout, fmt.Sprintf("%s[DEBUG]%s ", blue, reset), log.Ldate|log.Ltime)
)

func Info(msg string) {
	infoLogger.Output(2, msg)
}

func Warn(msg string) {
	warnLogger.Output(2, msg)
}

func Error(msg string) {
	errorLogger.Output(2, msg)
}

func Debug(msg string) {
	debugLogger.Output(2, msg)
}
