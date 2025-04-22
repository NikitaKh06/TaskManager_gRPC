package config

import (
	"log"
	"os"
)

var LogFile *os.File

func ConfigureLogger() error {
	logDirPath := "/app/logs"
	logFilePath := logDirPath + "/task-manager-api-logs.txt"

	err := os.MkdirAll(logDirPath, 0777)
	if err != nil {
		return err
	}

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	LogFile = logFile

	log.SetOutput(LogFile)

	log.Println("Logger configured")

	return nil
}
