package controller

import (
	"net/http"
	"strconv"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/service"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
}

func NewUser(service service.UserService) *userController {
	return &userController{service}
}

func (con *userController) UpdateUser(c *gin.Context) {
	var userReq app.UserRegister
	ID, _ := strconv.Atoi(c.Param("id"))
	// Check error if user requests do not pass the rules
	if err := c.ShouldBindJSON(&userReq); err != nil {
		err, httpCode := helpers.UserValidation(err)
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	// Hash password using bcrypt
	hashedPassword, err := helpers.Hash(userReq.Password)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	userReq.Password = hashedPassword

	user, err, httpCode := con.service.UpdateUser(&userReq, ID)

	// Check error while insert to database
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	// Formating response and return to the user
	response := helpers.UserResponseSuccess(user, "updated")
	c.JSON(http.StatusOK, response)
}

func (con *userController) DeleteUser(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	user, err, httpCode := con.service.DeleteUser(ID)

	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}
	// Formating response and return to the user
	response := helpers.UserResponseSuccess(user, "deleted")
	c.JSON(http.StatusOK, response)
	return
}
