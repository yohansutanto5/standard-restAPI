package main

import (
	"app/handler"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Rest API
func setupRoutes() *gin.Engine {
	r := gin.New()
	// Setup Middleware
	r.Use(middleware, gin.Recovery())
	// Setup Admin Endpoint
	var secrets = gin.H{
		"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
		"austin": gin.H{"email": "austin@example.com", "phone": "666"},
		"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
	}

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// Initiate all services and dependency
	userService := service.NewUserService(database)
	UserProfileService := service.NewUserProfileService(database)
	// Define The route Path
	// System API
	r.GET("/health", func(c *gin.Context) { handler.GetSystemHealth(c, database) })

	// Post Method
	r.POST("/user", func(c *gin.Context) { handler.InsertUser(c, userService) })
	r.POST("/userprofile", func(c *gin.Context) { handler.InsertUserProfile(c, UserProfileService) })

	// Get Method
	r.GET("/user", func(c *gin.Context) { handler.GetUser(c, userService) })
	r.GET("/userprofile", func(c *gin.Context) { handler.GetUserProfile(c, UserProfileService) })

	return r
}
