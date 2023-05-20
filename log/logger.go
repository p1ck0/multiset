package log

import (
	"io"
	"log"
	"os"
)

func New(debug bool) *Logger {

	logger := &Logger{
		infoLog:  log.New(os.Stdout, "| INFO | ", log.Ldate|log.Ltime|log.Lmsgprefix),
		errLog:   log.New(os.Stderr, "| ERROR | ", log.Ldate|log.Ltime|log.Lmsgprefix),
		debugLog: log.New(io.Discard, "| DEBUG | ", log.Ldate|log.Ltime|log.Lmsgprefix),
	}

	if debug {
		logger.debugLog.SetOutput(os.Stdout)
	}

	return logger
}

type Logger struct {
	infoLog  *log.Logger
	errLog   *log.Logger
	debugLog *log.Logger
}

func (l *Logger) Error(msg string) {
	l.errLog.Println(msg)
}

func (l *Logger) Info(msg string) {
	l.infoLog.Println(msg)
}

func (l *Logger) Debug(msg string) {
	l.debugLog.Println(msg)
}
