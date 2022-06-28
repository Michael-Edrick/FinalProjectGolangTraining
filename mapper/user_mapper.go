package mapper

import "FinalProject/entity"

func RegisterMapper(response entity.User) entity.UserRegister{
	var userRegister entity.UserRegister
	userRegister.RegisterAge = response.Age
	userRegister.RegisterEmail = response.Email
	userRegister.RegisterId = response.Id
	userRegister.RegisterUsername = response.Username
	return userRegister
}

func UpdateMapper(response entity.User)entity.UserUpdate{
	var UserUpdate entity.UserUpdate
	UserUpdate.Id = response.Id
	UserUpdate.Username = response.Username
	UserUpdate.Email = response.Email
	UserUpdate.Password = response.Password
	UserUpdate.Age = response.Age
	UserUpdate.Updated_at = response.Updated_at
	return UserUpdate
}