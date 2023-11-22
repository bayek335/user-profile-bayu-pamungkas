package model

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"gorm.io/gorm"
)

type UserModel interface {
	FindByEmail(user *app.User) (*app.User, error)
	CreateUser(userReq *app.User) (*app.User, error)
	UpdateUser(userReq *app.User, ID int) (*app.User, error)
	DeleteUser(ID int) (*app.User, error)
}

type userModel struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userModel {
	return &userModel{db}
}

func (m *userModel) FindByEmail(user *app.User) (*app.User, error) {

	err := m.db.Where("email=?", user.Email).First(&user).Error

	return user, err
}

func (m *userModel) CreateUser(user *app.User) (*app.User, error) {

	// Save the user data to database and return user data while no error
	// also return error if it is founded
	err := m.db.Create(&user).Error

	return user, err
}

func (m *userModel) UpdateUser(user *app.User, ID int) (*app.User, error) {
	var err error
	var User app.User
	// if just send request value is active
	// the upddate user account to actvated
	// if send all request value then update all user data
	if user.Username == "" {
		if user.IsActive {
			err = m.db.Model(&user).Where("id = ?", ID).Update("is_active", user.IsActive).Error
		}
	} else {
		res := m.db.Model(&User).Where("id = ?", ID).Updates(&user)
		// if there is not rows changed value
		// it mean user does not exist
		if res.RowsAffected < 1 {
			err := gorm.ErrRecordNotFound
			return user, err
		}
	}
	user.ID = ID
	return user, err
}

func (m *userModel) DeleteUser(ID int) (user *app.User, err error) {
	var User *app.User

	// check user is exist or not
	// delete user according the id as primary key if exist
	err = m.db.First(&User, ID).Error
	if err != nil {
		return user, err
	}

	err = m.db.Delete(&User, ID).Error
	return User, err
}
