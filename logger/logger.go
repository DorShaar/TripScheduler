package logging

import (
	"log"
	"os"
	"time"
)

type Logger struct {
	isInitialized bool
	fileName      string
}

func (logger *Logger) init() {
	if !logger.isInitialized {
		logger.fileName = time.Now().Format("2012-11-01T22:08:41")
		logger.isInitialized = true
	}
}

func (logger Logger) Log(msg string) {
	logFile, err := os.OpenFile(logger.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println(msg)
}
