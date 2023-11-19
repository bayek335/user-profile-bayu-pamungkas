package model

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"gorm.io/gorm"
)

type UserModel interface {
	CreateUser(userReq *app.User) (*app.User, error)
}

type userModel struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userModel {
	return &userModel{db}
}

func (m *userModel) CreateUser(user *app.User) (*app.User, error) {

	// Save the user data to database and return user data while no error
	// also return error if it is founded
	err := m.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}
