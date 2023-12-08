package db

import (
	"app/model"
)

type UserRepository struct{
	
}

func (d *DataStore) GetListUser() (Users []model.User, err error) {
	err = d.Db.Preload("Profile").Find(&Users).Error
	return
}

func (d *DataStore) InsertUser(User *model.User) (err error) {
	err = d.Db.Create(User).Error
	return
}

func (d *DataStore) DeleteUserByID(id int) (err error) {
	err = d.Db.Where("id = ?", id).Delete(&model.User{}).Error
	return
}

func (d *DataStore) UpdateUser(User *model.User) (err error) {
	err = d.Db.Save(&User).Error
	return
}
