package main

import (
	"app/handler"

	"github.com/gin-gonic/gin"
)

// Rest API
func setupRoutes() *gin.Engine {
	r := gin.New()
	// Setup Middleware
	r.Use(middleware, gin.Recovery())

	// Initiate all Handler and dependency
	userHandler := handler.NewUserHandler(database)
	UserProfilehandler := handler.NewUserProfileHandler(database)
	systemHander := handler.NewSystemHandler(database)

	// Define The route Path
	// ---- System API ---
	r.GET("/health", systemHander.GetSystemHealth)

	// ---- User API ---
	r.POST("/user", userHandler.Insert)
	r.GET("/user", userHandler.GetList)

	// ---- UserProfile API ---
	r.POST("/userprofile", UserProfilehandler.Insert)
	r.GET("/userprofile", UserProfilehandler.GetList)

	return r
}
