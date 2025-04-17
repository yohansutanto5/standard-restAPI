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

// Function Implementation

func (s UserServiceImpl) GetList() ([]model.User, *error.Error) {
	return s.db.GetListUser()
}

func (s *UserServiceImpl) Insert(User *model.User) *error.Error {
	return s.db.InsertUser(User)
}