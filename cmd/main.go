package main

import (
	"app/cmd/config"
	"app/db"
)

var database *db.DataStore

func init() {
	// setup Configuration
	var configuration config.Configuration = config.Load("test")

	// setup DB connection
	database = db.NewDatabase(configuration)
}

func main() {
	// Initiate UP SQL Migrations
	// If fail will execute down migrations then exit the application
	db.Migration()

	// Setup Gin Route
	r := setupRoutes()
	r.Run(":8078")
}
