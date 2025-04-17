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
	systemHander := handler.NewSystemHandler(database)

	UserProfilehandler := handler.NewUserProfileHandler(database)
		Userhandler:= handler.NewUserHandler(database)
 //{.NewHandler}

	// Define The route Path
	// ---- System API ---
	r.GET("/health", systemHander.GetSystemHealth)

	r.GET("/userprofile", UserProfilehandler.GetList)
	r.POST("/userprofile", UserProfilehandler.Insert)

	// ---- User API ---
		r.GET("/user", Userhandler.GetList)
	r.POST("/user", Userhandler.Insert)
	// ---- {.model} API ---
 //{.NewRoute}

	return r
}
