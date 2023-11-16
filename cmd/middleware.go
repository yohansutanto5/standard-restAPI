package main

import (
	"app/constanta"
	"app/model"
	"app/pkg/log"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	transactionID := generateTransactionID()
	c.Set("transactionID", transactionID)
	start := time.Now()
	// Incoming request logging
	incomingLog := model.CustomLog{
		Agent:         c.Request.UserAgent(),
		Method:        c.Request.Method,
		ClientIp:      c.ClientIP(),
		Path:          c.Request.URL.Path,
		TransactionID: transactionID,
		Status:        200,
		Code:          constanta.CodeOK,
		Message:       "Incoming Request",
	}
	log.Info(incomingLog)

	defer func() {
		if err := recover(); err != nil {
			// Handle the error, log it, and send an appropriate response.
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			// log.Error(fmt.Sprintf("%v", err))
		}
	}()

	// Process the request
	c.Next()

	end := time.Now()
	// Response logging
	responseLog := model.CustomLog{
		Agent:         c.Request.UserAgent(),
		Method:        c.Request.Method,
		ClientIp:      c.ClientIP(),
		Path:          c.Request.URL.Path,
		TransactionID: transactionID,
		Status:        c.Writer.Status(),
		Duration:      end.Sub(start),
		Code:          constanta.CodeOK,
		Message:       "Response Message please fix this and the code so that it is dynamic",
	}
	log.Info(responseLog)

}

func customLogFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

func generateTransactionID() int {
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
