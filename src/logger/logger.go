package logger

import (
	"log"
	"os"
)

type logger struct{}

// This makes sure that the logger type always implements the Logger interface
var _ Logger = (*logger)(nil)

var (
	// Info logger is to log Info type logs
	infoLogger *log.Logger
	// Warning logger is to log Warning type logs
	warningLogger *log.Logger
	// errorLogger is to log Error type logs
	errorLogger *log.Logger
)

// GetLogger initialises and returns an instance of a logger
func GetLogger() Logger {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	return &logger{}
}

// Info is for logging Info logs
func (l logger) Info(v ...interface{}) {
	infoLogger.Println(v...)
}

// Warn is for logging warnings
func (l logger) Warn(v ...interface{}) {
	warningLogger.Println(v...)
}

// Error is for logging errors
func (l logger) Error(v ...interface{}) {
	errorLogger.Println(v...)
}
