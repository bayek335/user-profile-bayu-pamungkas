package service

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/model"
)

type PhotoService interface {
	CreatePhoto(photo *app.Photos) (*app.Photos, error)
	GetPhoto(ID int) (*app.Photos, error)
	UpdatePhoto(photo *app.Photos, ID int) (*app.Photos, error)
	DeletePhoto(ID int) (*app.Photos, error)
}

type photoService struct {
	model model.PhotoModel
}

func NewPhoto(model model.PhotoModel) *photoService {
	return &photoService{model}
}

func (s *photoService) CreatePhoto(photo *app.Photos) (*app.Photos, error) {
	var Photo *app.Photos
	photo, err := s.model.CreatePhoto(Photo)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *photoService) GetPhoto(ID int) (*app.Photos, error) {
	photo, err := s.model.GetPhoto(ID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *photoService) UpdatePhoto(photo *app.Photos, ID int) (*app.Photos, error) {
	var Photo *app.Photos
	photo, err := s.model.UpdatePhoto(Photo, ID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *photoService) DeletePhoto(ID int) (*app.Photos, error) {
	photo, err := s.model.DeletePhoto(ID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}
