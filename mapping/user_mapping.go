package mapping

import "FinalProject/entity"

func RegisterMapping(response entity.User) entity.UserRegister{
	var userRegister entity.UserRegister
	userRegister.RegisterAge = response.Age
	userRegister.RegisterEmail = response.Email
	userRegister.RegisterId = response.Id
	userRegister.RegisterUsername = response.Username
	return userRegister
}