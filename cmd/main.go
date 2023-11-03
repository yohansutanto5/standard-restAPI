package main

import (
	"app/cmd/config"
	"app/db"

	"github.com/gin-gonic/gin"
)

var database *db.Database

func init() {
	// setup Configuration
	var configuration config.Configuration = config.Load("dev")

	// setup DB connection
	database = db.NewDatabase(configuration)
}

func main() {
	// Setup Gin Route
	r := setupRoutes()
	r.Use(middleware, gin.LoggerWithFormatter(customLogFormatter), gin.Recovery())
	r.Run(":8080")
}
