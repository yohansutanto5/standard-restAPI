package db

import (
	"app/model"
	"app/pkg/error"
)

func (d *DataStore) GetListUser() (Users []model.User, err *error.Error) {
	err = &error.Error{}
	e := d.Db.Preload("Profile").Find(&Users).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) InsertUser(User *model.User) (err *error.Error) {
	err = &error.Error{}
	e := d.Db.Create(User).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) DeleteUserByID(id int) (err *error.Error) {
	err = &error.Error{}
	e := d.Db.Where("id = ?", id).Delete(&model.User{}).Error
	err.ParseMysqlError(e)
	return
}

func (d *DataStore) UpdateUser(User *model.User) (err *error.Error) {
	err = &error.Error{}
	e := d.Db.Save(&User).Error
	err.ParseMysqlError(e)
	return
}
