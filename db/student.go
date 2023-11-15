package db

import (
	"app/model"
)

func (d *Database) GetListStudent() ([]model.Student, error) {
	var students []model.Student
	res := d.Db.Find(&students)
	if res.Error != nil {
		return nil, res.Error
	}
	return students, nil
}

func (d *Database) InsertStudent(student *model.Student) error {
	return d.Db.Create(student).Error
}

func (d *Database) DeleteStudentByID(id int) error {
	err := d.Db.Where("id = ?", id).Delete(&model.Student{}).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (d *Database) UpdateStudent(student *model.Student) error {
	err := d.Db.Save(&student).Error
	if err != nil {
		return err
	}
	return nil
}
