package logging

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type Logger struct {
	isInitialized bool
	fileName      string

	LogsDirectory string
}

func (logger *Logger) Init() {
	if !logger.isInitialized {
		// Create logs directory.
		if logger.LogsDirectory == "" {
			logger.LogsDirectory = "Logs"
		}

		os.Mkdir(logger.LogsDirectory, os.ModeDir)

		// Create logs filename
		t := time.Now()
		fileName := fmt.Sprintf("%d-%02d-%02dT%02d-%02d-%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second()) + ".log"

		logger.fileName = path.Join(logger.LogsDirectory, fileName)
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
