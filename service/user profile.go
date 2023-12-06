package service

import (
	"app/db"
	"app/model"
	"app/pkg/error"
)

// UserProfileService defines the interface for managing UserProfiles.
type UserProfileService interface {
	Insert(UserProfile *model.UserProfile) *error.Error
	Update(data *model.UserProfile) *error.Error
	DeleteByID(id int) *error.Error
	GetList() ([]model.UserProfile, *error.Error)
}

type UserProfileServiceImpl struct {
	db *db.DataStore
}

func NewUserProfileService(db *db.DataStore) UserProfileService {
	return &UserProfileServiceImpl{db: db}
}

// Function Implementation

func (s UserProfileServiceImpl) GetList() ([]model.UserProfile, *error.Error) {
	return s.db.GetListUserProfile()
}

func (s *UserProfileServiceImpl) DeleteByID(id int) *error.Error {
	return s.db.DeleteUserProfileByID(id)
}

func (s *UserProfileServiceImpl) Update(data *model.UserProfile) *error.Error {
	return s.db.UpdateUserProfile(data)
}

func (s *UserProfileServiceImpl) Insert(UserProfile *model.UserProfile) *error.Error {
	return s.db.InsertUserProfile(UserProfile)
}
