package main

import (
	"app/handler"
	"app/service"

	"github.com/gin-gonic/gin"
)

func setupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(middleware, gin.LoggerWithFormatter(customLogFormatter), gin.Recovery())
	// Initiate all services
	studentService := service.NewStudentService(database)

	// Define The route Path
	r.GET("/template", func(c *gin.Context) { handler.GetStudent(c, studentService) })
	r.POST("/template", handler.AddStudent)
	r.DELETE("/template/:id", handler.DeleteStudent)
	r.PUT("/template/:id", handler.UpdateStudent)
	return r
}