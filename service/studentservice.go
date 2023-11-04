package service

import (
	"app/db"
	"app/model"
)

func NewStudent(firstname, lastname string) model.Student {
	return model.Student{
		FirstName: firstname,
		LastName:  lastname,
	}
}

// StudentService defines the interface for managing students.
type StudentService interface {
	Create(student *model.Student) error
	GetByID(id int) (*model.Student, error)
	Update(data *model.Student) error
	DeleteByID(id int) error
	New(FirstName, LastName string, id int) model.Student
	GetList() ([]model.Student, error)
}

type StudentServiceImpl struct {
	db *db.Database
}

func NewStudentService(db *db.Database) StudentService {
	return &StudentServiceImpl{db: db}
}

// Function Implementation
func (s *StudentServiceImpl) GetByID(id int) (*model.Student, error) {
	// Implementation for fetching a student by ID from the database
	student := &model.Student{}
	if err := s.db.Db.Delete(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentServiceImpl) DeleteByID(id int) error {
	// Implementation for fetching a student by ID from the database
	if err := s.db.Db.Where("id = ?", id).Delete(&model.Student{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudentServiceImpl) Update(data *model.Student) error {
	if err := s.db.Db.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudentServiceImpl) Create(student *model.Student) error {
	return s.db.InsertStudent(student)
}

func (s *StudentServiceImpl) New(FirstName, LastName string, id int) model.Student {
	var st model.Student
	st.FirstName = FirstName

	st.LastName = LastName
	st.ID = id
	return st
}

func (s StudentServiceImpl) GetList() ([]model.Student, error) {
	students, err := s.db.GetListStudent()
	if err != nil {
		return nil, err
	}
	return students, nil
}
