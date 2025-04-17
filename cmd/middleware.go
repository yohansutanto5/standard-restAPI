package main

import (
	"app/constanta"
	"app/model"
	"app/pkg/error"
	"app/pkg/log"
	"bytes"
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	transactionID := generateTransactionID()
	c.Set("transactionID", transactionID)
	start := time.Now()
	defer func() {
		if err := recover(); err != nil {
			// Handle the error, log it, and send an appropriate response.
			c.JSON(500, gin.H{"error": "Internal Server Error"})
			log.Warning(transactionID, fmt.Sprintf("%v", err), printStack(5))
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
	if c.Writer.Status() <= 404 {
		responseLog.Code = constanta.CodeOK
		responseLog.Message = constanta.SuccessMessage
		log.Info(responseLog)
	} else {
		if lastError := c.Errors.Last(); lastError != nil {
			if v := lastError.Meta; v != nil {
				if e, ok := v.(*error.Error); ok {
					responseLog.Code = e.Response.Code
					responseLog.Message = e.Response.Message
					responseLog.Data = e.Response.Details
					log.Error(responseLog)
				}
			}
		} 
	}

}

func generateTransactionID() int {
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
}

func printStack(depth int) string {
	stackTrace := debug.Stack()
	lines := make([]string, 0)

	// Split the stack trace into lines
	for _, line := range bytes.Split(stackTrace, []byte{'\n'}) {
		lines = append(lines, string(line))
	}

	// Build the stack trace string up to the specified depth
	var buffer bytes.Buffer
	for i := 0; i < depth && i < len(lines); i++ {
		buffer.WriteString(lines[i])
		buffer.WriteString("\n")
	}
	return buffer.String()
}
