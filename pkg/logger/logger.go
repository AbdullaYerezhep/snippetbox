package logger

import (
	"log"
	"os"
)

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func New() Logger {
	return Logger{
		ErrorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime),
		InfoLog:  log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
