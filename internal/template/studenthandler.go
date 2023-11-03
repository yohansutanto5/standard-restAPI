package template

import (
	"app/constanta"
	"app/pkg/log"
	"app/pkg/util"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetStudent(c *gin.Context, db *gorm.DB) {
	var student = service.NewStudentService(db)
	// To DO handle filter and search
	c.JSON(http.StatusOK, student.GetList())
}

func DeleteStudent(c *gin.Context, id int, db *gorm.DB) {
	var student = service.NewStudentService(db)
	// To DO handle filter and search
	c.JSON(http.StatusOK, student.DeleteByID(id))
}

func AddStudent(c *gin.Context, data AddStudentIn, db *gorm.DB) {
	var student = service.NewStudentService(db)
	var newStudent service.Student
	newStudent.FirstName = data.FirstName
	newStudent.LastName = data.LastName
	err := student.CreateStudent(&newStudent)

	if err != nil {
		log.Error(util.GetTransactionID(c), err.Error(), nil)
		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
	}
	c.JSON(http.StatusOK, constanta.SuccessMessage)
}

func UpdateStudent(c *gin.Context, data AddStudentIn, db *gorm.DB) {
	var student = service.NewStudentService(db)
	var newStudent service.Student
	newStudent.FirstName = data.FirstName
	newStudent.LastName = data.LastName

	err := student.Update(&newStudent)
	if err != nil {
		log.Error(util.GetTransactionID(c), err.Error(), nil)
		c.JSON(http.StatusInternalServerError, constanta.InternalServerErrorMessage)
	}
	c.JSON(http.StatusOK, constanta.SuccessMessage)
}
