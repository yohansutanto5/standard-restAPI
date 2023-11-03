package testtemplate

import (
	"app/model"
	"app/pkg/log"
	"app/service"
	"fmt"
	"testing"
)

func TestMigrationTemplate(t *testing.T) {
	err := dbg.AutoMigrate(&service.Department{}, &service.Teacher{}, &service.Enrollment{}, &service.Student{}, &service.Course{})
	if err != nil {
		fmt.Println(err.Error())
		t.Failed()
	}
}

func TestCreateTemplate(t *testing.T) {
	var student model.Student
	student.FirstName = "halo3"
	student.LastName = "babi3"

	student.Insert(dbg)
}

func TestGetListTemplate(t *testing.T) {
	var student = service.NewStudentService(dbg)
	students := student.GetList()
	log.PrintStruct(students)
	if len(students) < 2 {
		t.FailNow()
	}
}

// func TestGormRelationCreateFinal(t *testing.T) {
// 	res := dbg.Create(&Cart{
// 		ID: 1,
// 		Product: []Product{
// 			{ID: 1, Name: "hy"}, {ID: 2, Name: "lpl"},
// 		},
// 		Total: 123,
// 	})
// 	if res.Error != nil {
// 		t.FailNow()
// 		fmt.Println(res.Error.Error())
// 	}
// }
