package utils

import (
	"log"
	"os"
)

type logger struct {
	info  *log.Logger
	debug *log.Logger
	err   *log.Logger
}

type Logger interface {
	Info(s string)
	Debug(s string)
	Error(s string)
}

func (l *logger) Info(s string) {
	l.info.Println(s)
}

func (l *logger) Debug(s string) {
	l.debug.Println(s)
}

func (l *logger) Error(s string) {
	l.err.Println("\033[0;31m"+s)
}

func CreateLogger() Logger {
	return &logger{
		info: log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate),
		debug: log.New(os.Stdout, "DEBUG: ", log.Ltime|log.Ldate),
		err: log.New(os.Stdout, "ERROR: ", log.Ltime|log.Ldate),
	}
}