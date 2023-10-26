package db

import "time"

type Student struct {
	ID          uint `gorm:"primaryKey"`
	FirstName   string
	LastName    string
	Enrollments []Enrollment // One-to-Many: One student can be enrolled in multiple courses
}

type Course struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	TeacherID uint      // Many-to-One: Many courses are taught by one teacher
	Students  []Student `gorm:"many2many:enrollments"` // Many-to-Many: Many students can enroll in many courses
}

type Teacher struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Courses      []Course // One-to-Many: One teacher can teach multiple courses
	DepartmentID uint
}

type Enrollment struct {
	ID             uint `gorm:"primaryKey"`
	StudentID      uint
	CourseID       uint
	EnrollmentDate time.Time
	Student        Student `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-One: Each enrollment is associated with one student
	Course         Course  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // One-to-One: Each enrollment is associated with one course
}

type Department struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Teachers []Teacher // One-to-Many: One department can have multiple teachers
}
