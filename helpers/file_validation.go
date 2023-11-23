package helpers

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidationFile(c *gin.Context, form string, ID int) (*multipart.FileHeader, string, error) {

	// get file and check if image has been sent
	file, err := c.FormFile("photo")
	if err != nil {
		err := errors.New("Field photo is required!")
		return file, "", err
	}

	fileSize := file.Size
	fileName := file.Filename
	fileExt := filepath.Ext(fileName)

	if !checkFileExt(fileExt) {
		err := errors.New("File must be type of image '.jpg', '.jpeg', '.png' !")
		return file, "", err
	}

	if fileSize > 2048000 {
		err := errors.New("Maximum size of photo is 2048Kb!")
		return file, "", err
	}

	newFileName := GenerateRandomString(15) + strconv.Itoa(ID)
	fileUrl := "public/images/" + newFileName + fileExt

	return file, fileUrl, nil
}

func checkFileExt(fileExt string) bool {
	if fileExt == ".jpg" || fileExt == ".jpeg" || fileExt == ".png" {
		return true
	}
	return false
}
