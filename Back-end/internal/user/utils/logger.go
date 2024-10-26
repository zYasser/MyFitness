package utils

import (
	"log"
	"os"
)

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}

func GetLogger() *Logger{
	 return &Logger{InfoLog: log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile), ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)}

}


