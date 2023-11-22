package service

import (
	"fmt"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/model"
)

type UserService interface {
	FindByEmail(userLogin *app.UserLogin) (*app.User, error)
	CreateUser(userReq *app.UserRegister) (*app.User, error, int)
	UpdateUser(userReq *app.UserRegister, ID int) (*app.User, error, int)
	DeleteUser(ID int) (*app.User, error, int)
}

type userService struct {
	model model.UserModel
}

var httpCode int

func NewUser(model model.UserModel) *userService {
	return &userService{model}
}

func (s *userService) FindByEmail(userLogin *app.UserLogin) (*app.User, error) {

	user := &app.User{
		Email: userLogin.Email,
	}
	user, err := s.model.FindByEmail(user)

	if err != nil {
		err, httpCode = helpers.DatabaseErrorHandler(err, "user")
		return user, err
	}
	fmt.Println("df")
	return user, nil
}

func (s *userService) CreateUser(userReq *app.UserRegister) (*app.User, error, int) {

	// Parsing user request struct format to user struct format
	// before send to model in order to insert to database
	user := &app.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	user, err := s.model.CreateUser(user)

	if err != nil {
		err, httpCode = helpers.DatabaseErrorHandler(err, "user")
	}
	return user, err, httpCode
}

func (s *userService) UpdateUser(user *app.UserRegister, ID int) (*app.User, error, int) {
	User := &app.User{}
	var httpCode int
	if user.Username != "" {
		User.Username = user.Username
		User.Email = user.Email
		User.Password = user.Password
	} else {
		User.IsActive = user.IsActive
	}

	_, err := s.model.UpdateUser(User, ID)
	if err != nil {
		err, httpCode = helpers.DatabaseErrorHandler(err, "User")
		return User, err, httpCode
	}
	return User, nil, httpCode
}

func (s *userService) DeleteUser(ID int) (*app.User, error, int) {
	var httpCode int

	user, err := s.model.DeleteUser(ID)
	if err != nil {
		err, httpCode = helpers.DatabaseErrorHandler(err, "User")
		return user, err, httpCode
	}
	return user, nil, httpCode
}
