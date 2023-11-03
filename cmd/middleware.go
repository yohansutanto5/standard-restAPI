package main

import (
	"app/pkg/log"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	transactionID := generateTransactionID()
	c.Set("transactionID", transactionID)
	c.Set("db", database)
	log.Info(transactionID, "incoming request", nil)
	defer func() {
		if err := recover(); err != nil {
			// Handle the error, log it, and send an appropriate response.
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			log.Error(transactionID, "PANIC", err)
		}
	}()
	c.Next()
	// Authentication here
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
