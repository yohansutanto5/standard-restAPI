package main

import (
	"app/cmd/config"
	"app/db"
	"app/pkg/log"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// setup Configuration
	var configuration config.Configuration = config.Load("dev")
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// setup DB connection
	// dbconn, err := db.ConnectDB(configuration.Db.Username, configuration.Db.Password, configuration.Db.Host, configuration.Db.Database)
	dbconn,err := db.GormInit(configuration)
	if err != nil {
		panic(err)
	}
	// Setup Route
	r := gin.New()
	r.Use(requestLoggerMiddleware(), middleware, gin.LoggerWithFormatter(customLogFormatter), gin.Recovery())
	config.LoadRoutes(r, dbconn)

	r.Run(":8080")
}

func middleware(c *gin.Context) {
	transactionID := generateTransactionID()
	c.Set("transactionID", transactionID)
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

func generateTransactionID() int {
	min := 100000
	max := 999999
	return rand.Intn(max-min+1) + min
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

func requestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request details
		latency := time.Since(time.Now())
		fmt.Printf("[%s] [%s] [%s] [%s] [%s] [%d] [%s] [%d] [%s]\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.Proto,
			c.Writer.Status(),
			latency,
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)
		// Pass the request to the next middleware or route handler
		c.Next()
	}
}
