package helpers

import "github.com/bayek335/task-5-pbi-btpns-bayu-pamungkas/app"

// Formating respon api while store data or validation request fail
func UserResponseFail(err error) *app.UserResponseFail {

	userResFail := &app.UserResponseFail{
		Success: false,
		Message: err.Error(),
	}
	return userResFail

}

// Formating respon api while store data successed
func UserResponseSuccess(user *app.User, mes string) *app.UserResponseSuccess {

	userResSuc := &app.UserResponseSuccess{}
	userResSuc.Success = true
	userResSuc.Message = "User succesfully " + mes
	userResSuc.Data.ID = user.ID
	userResSuc.Data.Email = user.Email
	userResSuc.Data.Username = user.Username

	return userResSuc

}

// func UserLoginResponse(token string) *app.LoginResponse {
// 	loginRes := &app.LoginResponse{
// 		Success: true,
// 		Message: "sucess",
// 	}
// 	loginRes.Data.Token = token
// 	return loginRes
// }
