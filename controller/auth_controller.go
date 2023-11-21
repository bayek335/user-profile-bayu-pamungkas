package controller

import (
	"errors"
	"net/http"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type authController struct {
	service service.UserService
}

func NewAuth(service service.UserService) *authController {
	return &authController{service}
}

func (con *authController) GenerateCode(c *gin.Context) {
	verifyCode := helpers.GenerateCode()
	// Get user email from email session after loged in
	email := sessions.Default(c).Get("email")

	// Check user already activated
	result := con.checkIsActive(email.(string))
	if result {
		err := errors.New("User already activated!")
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusOK, response)
		return
	}

	// Hash verify code in order to set cookie as value
	str, _ := helpers.Hash(verifyCode)
	c.SetCookie("verify-code", str, 120, "/", "localhost", true, false)

	// Send verify code to user email
	if err := helpers.SendEmailVerifyCode(email.(string), verifyCode); err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Formating response and return to the user
	response := helpers.ResponseVerifyCodeSuccess("Verification code has been sent")

	c.JSON(http.StatusCreated, response)
}

func (con *authController) ValidateCode(c *gin.Context) {
	var verifyCode app.ValidateCode

	// Get user data from session
	IDStr := sessions.Default(c).Get("id")
	ID := IDStr.(int)
	email := sessions.Default(c).Get("email").(string)

	// Check is user already activated or not
	result := con.checkIsActive(email)
	if result {
		err := errors.New("User already activated!")
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusOK, response)
		return
	}

	err := c.ShouldBindJSON(&verifyCode)

	if err != nil {
		err, httpCode := helpers.UserValidation(err)
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	// Check is exist or not verify code cookie
	cookie, err := c.Cookie("verify-code")
	if err != nil {
		err = errors.New("Verification code is expired generate again!")
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// Compare the verification code in the cookie with user verification code request
	if err = bcrypt.CompareHashAndPassword([]byte(cookie), []byte(verifyCode.Code)); err != nil {
		err = errors.New("Verification code is wrong!")
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// Update user is active to be true in database
	userReq := &app.UserRegister{
		IsActive: true,
	}
	_, err, _ = con.service.UpdateUser(userReq, ID)

	// Formating response and return to the user
	response := helpers.ResponseVerifyCodeSuccess("Your account has been verified")

	c.JSON(http.StatusOK, response)
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

	// Hash password using bcrypt
	hashedPassword, err := helpers.Hash(userReq.Password)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	userReq.Password = hashedPassword

	// Call the service of user
	user, err, httpCode := con.service.CreateUser(&userReq)

	// Check error while insert to database
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	// Generate verifycation code & create cookie

	// Formating response and return to the user
	response := helpers.UserResponseSuccess(user, "created")
	c.JSON(http.StatusCreated, response)
}

func (con *authController) Login(c *gin.Context) {
	var userReq app.UserLogin

	if err := c.ShouldBindJSON(&userReq); err != nil {
		// Check error if user requests do not pass the rules
		err, httpCode := helpers.UserValidation(err)
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	// Find user by email exist or not
	user, err := con.service.FindByEmail(&userReq)
	if err != nil {
		err = errors.New("Wrong password or email!")
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Compare is match password db and password rq
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password)); err != nil {
		err = errors.New("Wrong password or email!")
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Create session
	helpers.SetUserLogin(c, user.ID, user.Username, user.Email)

	// Generate JWT token
	err = helpers.GenerateJwtOAuth(c, user.Email)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// Formating response and return to the user
	response := helpers.UserResponseSuccess(user, "loged in")
	c.JSON(http.StatusOK, response)
}

func (con *authController) LogOut(c *gin.Context) {
	// Clear all session user
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	var User app.User

	response := helpers.UserResponseSuccess(&User, "loged out")
	// Delete jwt token from cookie
	c.SetCookie("token", "", -1, "/", "localhost", true, false)
	c.JSON(http.StatusOK, response)
}

// Function to check wheter the user account it is already verified
func (con *authController) checkIsActive(email string) bool {
	userLogin := &app.UserLogin{
		Email: email,
	}
	user, _ := con.service.FindByEmail(userLogin)
	return user.IsActive
}
