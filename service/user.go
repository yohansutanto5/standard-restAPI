package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
)

// UserService defines the interface for managing Users.
type UserService interface {
	Insert(User *model.User) *error.Error
	GetList() ([]model.User, *error.Error)
}

type UserServiceImpl struct {
	db *db.DataStore
}

func NewUserService(db *db.DataStore) UserService {
	return &UserServiceImpl{db: db}
}

func (s UserServiceImpl) GetList() ([]model.User, *error.Error) {
	err := &error.Error{}
	result, e := s.db.GetListUser()
	if e != nil {
		err.ParseMysqlError(e)
		return nil, err
	} else {
		return result, nil
	}

}

func (s *UserServiceImpl) Insert(User *model.User) *error.Error {
	err := &error.Error{}
	e := s.db.InsertUser(User)
	if e != nil {
		err.ParseMysqlError(e)
		return err
	}
	return nil
}
