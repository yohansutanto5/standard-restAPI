package util

import (
	"time"

	"github.com/gin-gonic/gin"
)

func EchoTest() string {
	return "Success"
}

func DateConvert() time.Time {
	return Date()
}

func GetTransactionID(c *gin.Context) int {
	transactionID, _ := c.Get("transactionID")
	return ConvertToInt(transactionID)
}
