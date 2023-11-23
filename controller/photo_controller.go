package controller

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type photoController struct {
	service service.PhotoService
}

func NewPhoto(service service.PhotoService) *photoController {
	return &photoController{service}
}

func errorMessage(mess string) *app.UserResponseFail {
	err := errors.New(mess)
	response := helpers.UserResponseFail(err)

	return response
}

func (con *photoController) CreatePhoto(c *gin.Context) {
	var photoReq *app.PhotoRequest

	user_ID := sessions.Default(c).Get("id").(int)

	err := c.Bind(&photoReq)
	// Check error if user requests do not pass the rules
	if err != nil {
		err, httpCode := helpers.UserValidation(err)
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	file, fileUrl, err := helpers.ValidationFile(c, "photo", user_ID)
	if err != nil {
		response := errorMessage(err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// // get file and check if image has been sent
	// file, err := c.FormFile("photo")
	// if err != nil {
	// 	response := errorMessage("Field photo is required!")
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // get filename, filesize, and file extention
	// fileSize := file.Size
	// fileName := file.Filename
	// fileExt := filepath.Ext(fileName)

	// // if filesize more then 2Mb
	// if fileSize > 2048000 {
	// 	response := errorMessage("Maximum size of photo is 2048Kb!")
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	// // if file is not photo such as jpg, jpeg, png
	// if !helpers.CheckFileExt(fileExt) {
	// 	response := errorMessage("File must be type of image '.jpg', '.jpeg', '.png' !")
	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }
	// // create new filename
	// newFileName := helpers.GenerateRandomString(15) + strconv.Itoa(ID)
	// fileUrl := "public/images/" + newFileName + fileExt

	// save file image to specific folder with new name
	err = c.SaveUploadedFile(file, fileUrl)
	// if err while saving file to static file
	if err != nil {
		response := errorMessage("Internal server error!")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// initiate new object photos
	photo := &app.Photos{
		Title:    photoReq.Title,
		Caption:  photoReq.Caption,
		PhotoUrl: "localhost:8080/" + fileUrl,
		UserId:   user_ID,
	}

	// send object to service
	photo, err = con.service.CreatePhoto(photo)
	if err != nil {
		// if error while saving to database, delete file
		os.Remove(fileUrl)
		response := errorMessage("Internal server error!")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// success
	response := helpers.PhotoResponseSuccess(photo, "created")
	c.JSON(http.StatusOK, response)

}

func (con *photoController) GetPhotos(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	// get all photos by user id and return as slice
	photo, err := con.service.GetPhotos(ID)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusBadRequest, response)
	}

	// Formating response and return to the user
	response := helpers.PhotoResponseAll(photo, "taken")
	c.JSON(http.StatusOK, response)
	return

}

func (con *photoController) UpdatePhoto(c *gin.Context) {
	var photoReq *app.PhotoRequest
	ID, _ := strconv.Atoi(c.Param("id"))

	user_ID := sessions.Default(c).Get("id").(int)

	err := c.Bind(&photoReq)
	// Check error if user requests do not pass the rules
	if err != nil {
		err, httpCode := helpers.UserValidation(err)
		response := helpers.UserResponseFail(err)
		c.JSON(httpCode, response)
		return
	}

	file, fileUrl, err := helpers.ValidationFile(c, "photo", user_ID)
	if err != nil {
		response := errorMessage(err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.SaveUploadedFile(file, fileUrl)
	if err != nil {
		response := errorMessage("Internal server error!")
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	photo := &app.Photos{
		Title:    photoReq.Title,
		Caption:  photoReq.Caption,
		PhotoUrl: "localhost:8080/" + fileUrl,
	}

	// if update fail delete file
	updPhoto, oldFileUrl, e := con.service.UpdatePhoto(photo, ID, user_ID)
	if e != nil {
		os.Remove(fileUrl)
		err, httpCode := helpers.DatabaseErrorHandler(e, "Photos")
		response := errorMessage(err.Error())
		c.JSON(httpCode, response)
		return
	}
	// if update success delete old file
	// because path name from db contain hostport, trim the hostport
	os.Remove(strings.TrimLeft(oldFileUrl, "localhost:8080/"))

	// return success
	response := helpers.PhotoResponseSuccess(updPhoto, "updated")
	c.JSON(http.StatusOK, response)

}

func (con *photoController) DeletePhoto(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))
	user_ID := sessions.Default(c).Get("id").(int)

	// delete photo according the user and photo id
	photo, fileUrl, err := con.service.DeletePhoto(ID, user_ID)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusNotFound, response)
		return
	}

	os.Remove(strings.Trim(fileUrl, "localhost:8080/"))

	// return success
	response := helpers.PhotoResponseSuccess(photo, "deleted")
	c.JSON(http.StatusOK, response)

}
