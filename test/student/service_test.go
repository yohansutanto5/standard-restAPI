package student_test

import (
	"app/model"
	"app/pkg/log"
	"fmt"
	"testing"
)

func TestMigrationTemplate(t *testing.T) {
	err := dbg.AutoMigrate(&model.Student{})
	if err != nil {
		fmt.Println(err.Error())
		t.Failed()
	}
}

func TestCreateTemplate(t *testing.T) {
	var student model.Student
	student.FirstName = "halo3"
	student.LastName = "babi3"

	studenService.Create(&student)
}

func TestGetListTemplate(t *testing.T) {
	students, _ := studenService.GetList()
	log.PrintStruct(students)
	if len(students) < 2 {
		t.FailNow()
	}
}
