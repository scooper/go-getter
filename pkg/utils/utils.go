package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	l.err.Println(s)
}

func CreateLogger() Logger {
	return &logger{
		info: log.New(os.Stdout, "INFO: ", log.Ltime|log.Ldate),
		debug: log.New(os.Stdout, "DEBUG: ", log.Ltime|log.Ldate),
		err: log.New(os.Stdout, "ERROR: ", log.Ltime|log.Ldate),
	}
}

func GetSetting(path string) (*os.File, error) {
	return getFile("settings", path)
}

func GetTemplate(path string) (*os.File, error) {
	return getFile("templates", path)
}

func getFile(base string, fpath string) (*os.File, error) {
	// look in directory of currently running process
	path, perr := filepath.Abs(fmt.Sprintf("./%s/%s", base, fpath))
	if perr != nil {
		return nil, perr
	}

	file, ferr := os.Open(path)
	if ferr != nil {
		return nil, ferr
	}

	return file, nil
}