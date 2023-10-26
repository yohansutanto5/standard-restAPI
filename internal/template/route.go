package template

import (
	"app/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddStudentIn struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
}

func Templatendpoint(r *gin.Engine, db *gorm.DB) *gin.Engine {
	r.GET("/template", func(c *gin.Context) {
		GetStudent(c, db)
	})
	r.POST("/template", func(c *gin.Context) {
		var data AddStudentIn
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		AddStudent(c, data, db)
	})
	r.DELETE("/template/:id", func(c *gin.Context) {
		DeleteStudent(c, util.ConvertToInt(c.Param("id")), db)
	})
	return r
}
