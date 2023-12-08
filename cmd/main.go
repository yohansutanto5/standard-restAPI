package main

import (
	"app/cmd/config"
	"app/db"
	"app/model"
	"app/pkg/log"
	"os"
)

var database *db.DataStore
var configuration config.Configuration

func init() {
	// setup Configuration
	mode := os.Getenv("GIN_MODE")
	configfolder := os.Getenv("config")
	configuration = config.Load(mode, configfolder)
	log.System("Loaded Configuration Environment:" + mode + " from :" + configfolder)
	// setup DB connection
	database = db.NewDatabase(configuration)
	log.System("Connection to Database is Established")
}

func main() {
	// Initiate UP SQL Migrations
	// If fail will execute down migrations then exit the application
	// db.Migration(&configuration, false)
	if true {
		log.System("SQL Migration is Starting")
		err := database.Db.AutoMigrate(model.UserProfile{}, model.User{})
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	// Setup Gin Route
	log.System("Setting Up API Routes")
	r := setupRoutes()
	log.System("Application is running")
	r.Run(":8078")
}
