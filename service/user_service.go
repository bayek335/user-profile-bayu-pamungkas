package service

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/model"
)

type UserService interface {
	CreateUser(userReq *app.UserRegister) (*app.User, error, int)
}

type userService struct {
	model model.UserModel
}

func NewUser(model model.UserModel) *userService {
	return &userService{model}
}

func (s *userService) CreateUser(userReq *app.UserRegister) (*app.User, error, int) {
	var httpCode int
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
