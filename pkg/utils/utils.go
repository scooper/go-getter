package utils

import (
	"fmt"
	"io"
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
	SetFile(f *os.File)
}

func (l *logger) SetFile(f *os.File) {
	mw := io.MultiWriter(os.Stdout, f)
	l.info.SetOutput(mw)
	l.debug.SetOutput(mw)
	l.err.SetOutput(mw)
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
	return getFile("settings", path, false)
}

func GetTemplate(path string) (*os.File, error) {
	return getFile("templates", path, false)
}

func GetLogs(path string) (*os.File, error) {
	return getFile("logs", path, true)
}

func getFile(base string, fpath string, create bool) (*os.File, error) {
	// look in directory of currently running process
	path, perr := filepath.Abs(fmt.Sprintf("./%s/%s", base, fpath))
	if perr != nil {
		return nil, perr
	}

	flags := os.O_APPEND

	if create {
		flags = flags | os.O_CREATE
	}

	file, ferr := os.OpenFile(path, flags, 0777)
	if ferr != nil {
		return nil, ferr
	}

	return file, nil
}