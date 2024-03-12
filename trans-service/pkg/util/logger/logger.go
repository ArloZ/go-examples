package logger

import "log"

var logger *log.Logger = log.Default()

func Infof(format string, param ...any) {
	// TODO log info message
	logger.Printf(format, param...)
}
