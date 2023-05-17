package log

import (
	"log"
	"os"
)

func New() *Logger {
	return &Logger{
		infoLog: log.New(os.Stdout, "| INFO | ", log.Ldate|log.Ltime|log.Lmsgprefix),
		errLog:  log.New(os.Stderr, "| ERROR | ", log.Ldate|log.Ltime|log.Lmsgprefix),
	}
}

type Logger struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func (l *Logger) Error(msg string) {
	l.errLog.Println(msg)
}

func (l *Logger) Info(msg string) {
	l.infoLog.Println(msg)
}
