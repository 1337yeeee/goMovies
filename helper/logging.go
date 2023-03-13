package helper

import (
	"log"
	"os"
)

var logFileName string

func SetLogFileName(name string) {
	logFileName = name
}

func CreateLogger() (*log.Logger, *os.File, error) {
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
	return logger, file, nil
}

func CloseLogger(file *os.File) {
	if file != nil {
		file.Close()
	}
}
