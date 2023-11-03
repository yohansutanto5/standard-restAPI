package service

import (
	"time"

	"gorm.io/gorm"
)

func NewStudent(firstname, lastname string) Student {
	return Student{
		FirstName: firstname,
		LastName:  lastname,
	}
}

type Student struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	FirstName   string
	LastName    string
	Enrollments []Enrollment // One-to-Many: One student can be enrolled in multiple courses
}

// StudentService defines the interface for managing students.
type StudentService interface {
	CreateStudent(student *Student) error
	GetByID(id int) (*Student, error)
	Update(data *Student) error
	DeleteByID(id int) error
	New(FirstName, LastName string, id int) Student
	GetList() []Student
}

type StudentServiceImpl struct {
	db *gorm.DB
}

func NewStudentService(db *gorm.DB) StudentService {
	return &StudentServiceImpl{db: db}
}

// Function Implementation
func (s *StudentServiceImpl) GetByID(id int) (*Student, error) {
	// Implementation for fetching a student by ID from the database
	student := &Student{}
	if err := s.db.Delete(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (s *StudentServiceImpl) DeleteByID(id int) error {
	// Implementation for fetching a student by ID from the database
	if err := s.db.Where("id = ?", id).Delete(&Student{}).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudentServiceImpl) Update(data *Student) error {
	if err := s.db.Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (s *StudentServiceImpl) CreateStudent(student *Student) error {
	return s.db.Create(student).Error
}

func (s *StudentServiceImpl) New(FirstName, LastName string, id int) Student {
	var st Student
	st.FirstName = FirstName

	st.LastName = LastName
	st.ID = id
	return st
}

func (s StudentServiceImpl) GetList() []Student {
	var students []Student
	res := s.db.Find(&students)
	if res.Error != nil {
		return nil
	}
	return students
}

func (s Student) Insert(db *gorm.DB) error {
	db.Create(&s)
	return db.Error
}

func (s Student) Delete(db *gorm.DB) error {
	db.Delete(&s)
	return db.Error
}

type Course struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Title      string
	TeacherID  uint         // Many-to-One: Many courses are taught by one teacher
	Enrollment []Enrollment `gorm:"many2many:enrollments"` // Many-to-Many: Many students can enroll in many courses
}

type Teacher struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	Name         string
	Courses      []Course // One-to-Many: One teacher can teach multiple courses
	DepartmentID uint
}

type Enrollment struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	StudentID      uint
	CourseID       uint
	EnrollmentDate time.Time
	Student        Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-One: Each enrollment is associated with one student
	Course         Course  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-One: Each enrollment is associated with one course
}

type Department struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Teachers []Teacher // One-to-Many: One department can have multiple teachers
}
