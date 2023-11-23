package model

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"gorm.io/gorm"
)

type PhotoModel interface {
	CreatePhoto(photo *app.Photos) (*app.Photos, error)
	GetPhotos(ID int) ([]*app.Photos, error)
	UpdatePhoto(photo *app.Photos, ID, user_ID int) (*app.Photos, string, error)
	DeletePhoto(ID, user_ID int) (*app.Photos, string, error)
}

type photoModel struct {
	db *gorm.DB
}

func NewPhoto(db *gorm.DB) *photoModel {
	return &photoModel{db}
}

func (m *photoModel) CreatePhoto(photo *app.Photos) (*app.Photos, error) {

	err := m.db.Create(&photo).Error
	return photo, err
}
func (m *photoModel) GetPhotos(ID int) ([]*app.Photos, error) {
	var arrPhoto []*app.Photos
	err := m.db.Find(&arrPhoto).Error
	return arrPhoto, err
}
func (m *photoModel) UpdatePhoto(photo *app.Photos, ID, user_ID int) (*app.Photos, string, error) {
	var Photo *app.Photos
	err := m.db.Where("id =? AND user_id=?", ID, user_ID).First(&Photo).Error
	if err != nil {
		return photo, "", err
	}
	oldFileUrl := Photo.PhotoUrl
	err = m.db.Where("id = ?", ID).Updates(&photo).Error

	photo.ID = ID
	photo.UserId = Photo.UserId

	return photo, oldFileUrl, err
}
func (m *photoModel) DeletePhoto(ID, user_ID int) (*app.Photos, string, error) {
	var Photo *app.Photos
	result := m.db.Where("id =? AND user_id=?", ID, user_ID)
	if err := result.First(&Photo).Error; err != nil {
		return Photo, "", err
	}
	fileUrl := Photo.PhotoUrl
	err := result.Delete(&Photo).Error
	if err != nil {
		return Photo, "", err
	}
	return Photo, fileUrl, nil
}
