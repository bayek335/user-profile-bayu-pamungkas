package helpers

import "github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"

func PhotoResponseSuccess(photo *app.Photos, mes string) *app.PhotoResponseSuccess {

	photoRes := &app.PhotoResponseSuccess{}
	photoRes.Success = true
	photoRes.Message = "Photo succesfully " + mes
	photoRes.Data.ID = photo.ID
	photoRes.Data.Title = photo.Title
	photoRes.Data.Caption = photo.Caption
	photoRes.Data.PhotoUrl = photo.PhotoUrl
	photoRes.Data.UserId = photo.UserId

	return photoRes
}
func PhotoResponseAll(photos []*app.Photos, mes string) *app.PhotoResponseAll {

	photoRes := &app.PhotoResponseAll{}
	photoRes.Success = true
	photoRes.Message = "Photo succesfully " + mes
	photoRes.Data = photos

	return photoRes
}
