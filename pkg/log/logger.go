package log

import (
	"app/constanta"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type CustomLog struct {
	TransactionID int         `json:"transactionID"`
	Code          string      `json:"code"`
	Status        int         `json:"status"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
}

var log = logrus.New()

func init() {
	// Set log file name using the current date
	logFileName := "/home/yohan/standard-restAPI/app-" + time.Now().Format("2006-01-02") + ".log"

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
		"message":       customLog.Message,
		"data":          customLog.Data,
	}).Info() // Use Info as a generic log level; you can customize this as needed
}

func Debug(transactionID int, message string, data interface{}) {
	log.WithFields(logrus.Fields{
		"transactionID": transactionID,
		"message":       message,
		"data":          data,
		"status":        constanta.DebugStatus,
		"Code":          constanta.DebugCode,
	}).Debug()
}

func Error(transactionID int, message string, code string, data interface{}) {
	log.WithFields(logrus.Fields{
		"transactionID": transactionID,
		"message":       message,
		"data":          data,
		"status":        constanta.ErrorStatus,
		"Code":          code,
	}).Error()
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

func Fatal(message string) {
	customLog := CustomLog{
		TransactionID: 0,
		Level:         "FATAL",
		Message:       message,
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
