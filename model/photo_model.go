package model

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"gorm.io/gorm"
)

type PhotoModel interface {
	CreatePhoto(photo *app.Photos) (*app.Photos, error)
	GetPhoto(ID int) (*app.Photos, error)
	UpdatePhoto(photo *app.Photos, ID int) (*app.Photos, error)
	DeletePhoto(ID int) (*app.Photos, error)
}

type photoModel struct {
	db *gorm.DB
}

func NewPhoto(db *gorm.DB) *photoModel {
	return &photoModel{db}
}

func (m *photoModel) CreatePhoto(photo *app.Photos) (*app.Photos, error) {
	return photo, nil
}
func (m *photoModel) GetPhoto(ID int) (*app.Photos, error) {
	var photo *app.Photos
	return photo, nil
}
func (m *photoModel) UpdatePhoto(photo *app.Photos, ID int) (*app.Photos, error) {
	return photo, nil
}
func (m *photoModel) DeletePhoto(ID int) (*app.Photos, error) {
	var photo *app.Photos
	return photo, nil
}
