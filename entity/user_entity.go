package entity

import (
	"time"
)

type UserRepositoryInterface interface{
	UserRegisterRepository(newuser User)(User, error)
	UserLoginRepository(newlogin User)(string, error)
	UserUpdateRepository(updateuser User)(User, error)
	UserDeleteRepository(deleteuser User)(error)
	UserDeletePhotoRepository(deleteuser User)(error)
	UserDeleteCommentRepository(deleteuser User)(error)
	UserDeleteSocMedRepository(deleteuser User)(error)
	GetUserId(loginEmail User)(int, error)
}

type UserServiceInterface interface {
	RegisterService(user User)(User, error)
	LoginService(user User)(string, error)
	UserUpdateService(user User)(User, error)
	UserDeleteService(user User)(error)
}

type User struct{
	Id       int	`json:"id"`
	Username string `json:"username"`
	Email    string	`json:"email"`
	Password string `json:"password"`
	Age      int	`json:"age"`
	Created_at time.Time	`json:"-"`
	Updated_at time.Time	`json:"updated_at"`
}

type UserRegister struct{
	RegisterAge int `json:"age"`
	RegisterEmail string `json:"email"`
	RegisterId int `json:"id"`
	RegisterUsername string `json:"username"`
}

type Token struct{
	JwtToken string	`json:"token"`
}
