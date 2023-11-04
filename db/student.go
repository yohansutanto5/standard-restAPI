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
