package db

import "app/service"

func (d Database) GetListStudent() []service.Student {
	var students []service.Student
	res := d.db.Find(&students)
	if res.Error != nil {
		return nil
	}
	return students
}
