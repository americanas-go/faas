package kafka

import "github.com/americanas-go/log"

type Logger struct {
}

func (l *Logger) Printf(msg string, args ...interface{}) {
	log.Debugf(msg, args...)
}

type ErrorLogger struct {
}

func (l *ErrorLogger) Printf(msg string, args ...interface{}) {
	log.Errorf(msg, args...)
}
