package service

import (
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"
	"github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/model"
)

type PhotoService interface {
	CreatePhoto(photo *app.Photos) (*app.Photos, error)
	GetPhotos(ID int) ([]*app.Photos, error)
	UpdatePhoto(photo *app.Photos, ID, user_ID int) (*app.Photos, string, error)
	DeletePhoto(ID, user_ID int) (*app.Photos, string, error)
}

type photoService struct {
	model model.PhotoModel
}

func NewPhoto(model model.PhotoModel) *photoService {
	return &photoService{model}
}

func (s *photoService) CreatePhoto(photo *app.Photos) (*app.Photos, error) {
	photo, err := s.model.CreatePhoto(photo)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *photoService) GetPhotos(ID int) ([]*app.Photos, error) {
	photo, err := s.model.GetPhotos(ID)
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (s *photoService) UpdatePhoto(photo *app.Photos, ID, user_ID int) (*app.Photos, string, error) {
	photo, oldFileUrl, err := s.model.UpdatePhoto(photo, ID, user_ID)
	if err != nil {
		return photo, oldFileUrl, err
	}
	return photo, oldFileUrl, nil
}

func (s *photoService) DeletePhoto(ID, user_ID int) (*app.Photos, string, error) {
	photo, fileUrl, err := s.model.DeletePhoto(ID, user_ID)
	if err != nil {
		return photo, "", err
	}
	return photo, fileUrl, nil
}
