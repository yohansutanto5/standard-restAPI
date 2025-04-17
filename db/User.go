package db

import (
	"app/model"
	"app/pkg/error"
)


func (d *DataStore) GetListUser() (Users []model.User, err *error.Error) {
	e := d.Db.Find(&Users).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) InsertUser(User *model.User) (err *error.Error) {
	e := d.Db.Create(User).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) DeleteUserByID(id int) (err *error.Error) {
	e := d.Db.Where("id = ?", id).Delete(&model.User{}).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) UpdateUser(User *model.User) (err *error.Error) {
	e := d.Db.Save(&User).Error
	err = error.ParseMysqlError(e)
	return
}