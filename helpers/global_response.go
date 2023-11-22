package helpers

import "github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"

func ResponseVerifyCodeSuccess(mes string) *app.ResponseWithOutData {
	response := &app.ResponseWithOutData{
		Success: true,
		Message: mes,
	}

	return response
}
