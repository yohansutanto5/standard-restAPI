package db

import (
	"app/model"
	"app/pkg/error"
)


func (d *DataStore) GetListUserProfile() (UserProfiles []model.UserProfile, err *error.Error) {
	e := d.Db.Find(&UserProfiles).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) InsertUserProfile(UserProfile *model.UserProfile) (err *error.Error) {
	e := d.Db.Create(UserProfile).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) DeleteUserProfileByID(id int) (err *error.Error) {
	e := d.Db.Where("id = ?", id).Delete(&model.UserProfile{}).Error
	err = error.ParseMysqlError(e)
	return
}

func (d *DataStore) UpdateUserProfile(UserProfile *model.UserProfile) (err *error.Error) {
	e := d.Db.Save(&UserProfile).Error
	err = error.ParseMysqlError(e)
	return
}