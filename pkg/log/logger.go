package log

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomLog struct {
	TransactionID int         `json:"transactionID"`
	Level         string      `json:"level"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
}

var log = logrus.New()

func init() {
	// Set log file name using the current date
	logFileName := "app-" + time.Now().Format("2006-01-02") + ".log"

	// Open the log file for writing, create it if it doesn't exist
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = logFile
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	// Configure logrus to produce JSON logs
	log.SetFormatter(&logrus.JSONFormatter{})
}

func logWithFields(customLog CustomLog) {
	log.WithFields(logrus.Fields{
		"transactionID": customLog.TransactionID,
		"level":         customLog.Level,
		"message":       customLog.Message,
		"data":          customLog.Data,
	}).Info() // Use Info as a generic log level; you can customize this as needed
}

func Debug(transactionID int, message string, data interface{}) {
	customLog := CustomLog{
		TransactionID: transactionID,
		Level:         "DEBUG",
		Message:       message,
		Data:          data,
	}
	logWithFields(customLog)
}

func Error(transactionID int, message string, data interface{}) {
	customLog := CustomLog{
		TransactionID: transactionID,
		Level:         "ERROR",
		Message:       message,
		Data:          data,
	}
	logWithFields(customLog)
}

func Info(transactionID int, message string, data interface{}) {
	customLog := CustomLog{
		TransactionID: transactionID,
		Level:         "INFO",
		Message:       message,
		Data:          data,
	}
	logWithFields(customLog)
}

func Warning(transactionID int, message string, data interface{}) {
	customLog := CustomLog{
		TransactionID: transactionID,
		Level:         "warning",
		Message:       message,
		Data:          data,
	}
	logWithFields(customLog)
}

func PrintStruct(data interface{}) {
	customLog := CustomLog{
		Data: data,
	}
	logWithFields(customLog)
}
