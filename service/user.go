package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
)

// UserService defines the interface for managing Users.
type UserService interface {
	Insert(User *model.User) *error.Error
	Update(data *model.User) *error.Error
	DeleteByID(id int) *error.Error
	GetList() ([]model.User, *error.Error)
}

type UserServiceImpl struct {
	db *db.DataStore
}

func NewUserService(db *db.DataStore) UserService {
	return &UserServiceImpl{db: db}
}

func (s UserServiceImpl) GetList() ([]model.User, *error.Error) {
	return s.db.GetListUser()
}

func (s *UserServiceImpl) DeleteByID(id int) *error.Error {
	return s.db.DeleteUserByID(id)
}

func (s *UserServiceImpl) Update(data *model.User) *error.Error {
	return s.db.UpdateUser(data)
}

func (s *UserServiceImpl) Insert(User *model.User) *error.Error {
	return s.db.InsertUser(User)
}