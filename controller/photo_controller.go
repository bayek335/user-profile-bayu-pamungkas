package controller

import (
	"net/http"
	"strconv"

	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/helpers"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/model"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/service"
	"github.com/gin-gonic/gin"
)

type photoController struct {
	service service.PhotoService
}

func NewPhoto(model model.PhotoModel) *photoController {
	return &photoController{model}
}

func (con *photoController) CreatePhoto(c *gin.Context) {
	var Photo *app.Photos

	photo, err := con.service.CreatePhoto(Photo)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusBadRequest, response)
	}

	c.JSON(http.StatusOK, photo)

}

func (con *photoController) GetPhoto(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	photo, err := con.service.GetPhoto(ID)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusBadRequest, response)
	}

	c.JSON(http.StatusOK, photo)

}

func (con *photoController) UpdatePhoto(c *gin.Context) {
	var Photo *app.Photos
	ID, _ := strconv.Atoi(c.Param("id"))

	photo, err := con.service.UpdatePhoto(Photo, ID)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusBadRequest, response)
	}

	c.JSON(http.StatusOK, photo)

}

func (con *photoController) DeletePhoto(c *gin.Context) {
	ID, _ := strconv.Atoi(c.Param("id"))

	photo, err := con.service.DeletePhoto(ID)
	if err != nil {
		response := helpers.UserResponseFail(err)
		c.JSON(http.StatusBadRequest, response)
	}

	c.JSON(http.StatusOK, photo)

}
