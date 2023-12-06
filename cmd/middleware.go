package main

import (
	"app/constanta"
	"app/model"
	"app/pkg/error"
	"app/pkg/log"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// Group using gin.BasicAuth() middleware
// gin.Accounts is a shortcut for map[string]string

func middleware(c *gin.Context) {
	transactionID := generateTransactionID()
	c.Set("transactionID", transactionID)
	start := time.Now()
	defer func() {
		if err := recover(); err != nil {
			// Handle the error, log it, and send an appropriate response.
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			log.Warning(transactionID, fmt.Sprintf("%v", err), "Recover from Panic")
		}
	}()

	// Process the request
	c.Next()

	// Response logging
	responseLog := model.CustomLog{
		Agent:         c.Request.UserAgent(),
		Method:        c.Request.Method,
		ClientIp:      c.ClientIP(),
		Path:          c.Request.URL.Path,
		TransactionID: transactionID,
		Status:        c.Writer.Status(),
		Duration:      time.Duration(time.Since(start).Milliseconds()),
	}
	if c.Writer.Status() <= 400 {
		responseLog.Code = constanta.CodeOK
		responseLog.Message = constanta.SuccessMessage
		log.Info(responseLog)
	} else {
		v := c.Errors.Last().Meta //gin errr
		e, _ := v.(*error.Error)
		responseLog.Code = e.Response.Code
		responseLog.Message = e.Response.Message
		responseLog.Data = e.Response.Details
		log.Error(responseLog)
	}

}

func generateTransactionID() int {
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}
