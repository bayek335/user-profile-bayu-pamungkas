package controller

import (
	"net/http"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/service"
	"github.com/gin-gonic/gin"
)

type authController struct {
	service service.UserService
}

func NewAuth(service service.UserService) *authController {
	return &authController{service}
}

func (con *authController) Register(c *gin.Context) {

	var userReq app.UserRegister

	// Check error if user requests do not pass the rules
	if err := c.ShouldBindJSON(&userReq); err != nil {
		err, httpCode := helpers.UserValidation(err)
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	// Call the service of user
	user, err, httpCode := con.service.CreateUser(&userReq)

	// Check error while insert to database
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}
	// Formating response and return to the user
	response := helpers.UserResponseSuccess(user, "created")
	c.JSON(http.StatusCreated, response)
}
