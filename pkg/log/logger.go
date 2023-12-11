package log

import (
	"app/model"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	if os.Getenv("MODE") == "DEBUG" {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}
	// Set log file name using the current date
	logFileName := "../app-" + time.Now().Format("2006-01-02") + ".log"

	// Open the log file for writing, create it if it doesn't exist
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = logFile
	} else {
		log.SetOutput(os.Stdout)
		log.Info("Failed to log to file, using default stderr")
		return
	}

	// Configure logrus to produce JSON logs
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(io.MultiWriter(os.Stdout, log.Out))
}

func Debug(transactionID int, message string, data interface{}) {
	log.WithFields(logrus.Fields{
		"transactionID": transactionID,
		"data":          data,
	}).Debug(message)
}

func Error(param model.CustomLog) {
	log.WithFields(logrus.Fields{
		"transactionID": param.TransactionID,
		"status":        param.Status,
		"Code":          param.Code,
		"clientIp":      param.ClientIp,
		"method":        param.Method,
		"path":          param.Path,
		"agent":         param.Agent,
		"duration":      param.Duration,
	}).Error(param.Message)
}

func Info(param model.CustomLog) {
	log.WithFields(logrus.Fields{
		"transactionID": param.TransactionID,
		"status":        param.Status,
		"Code":          param.Code,
		"clientIp":      param.ClientIp,
		"method":        param.Method,
		"path":          param.Path,
		"agent":         param.Agent,
		"duration":      param.Duration,
	}).Info(param.Message)
}

func Fatal(message string) {
	log.Fatal(message)
}

func System(message string) {
	log.Info(message)
}

func Warning(transactionID int, message string, data interface{}) {
	log.WithFields(logrus.Fields{
		"transactionID": transactionID,
		"data":          data,
	}).Warning(message)
}
