package utils

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime | log.Lshortfile ), ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)}
	})
	return instance
}

