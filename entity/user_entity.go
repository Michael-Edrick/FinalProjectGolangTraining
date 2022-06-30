package entity

import (
	"time"
)

type UserServiceInterface interface {
	UserRegisterService(user *User) (*User, error)
	UserLoginService(user *User) (string, error)
	UserUpdateService(user *User) (*User, error)
	UserDeleteService(user *User) error
}

type UserRepositoryInterface interface {
	UserRegisterRepository(newuser *User) (*User, error)
	UserLoginRepository(newlogin *User) (*User, error)
	UserUpdateRepository(updateuser *User) (*User, error)
	UserDeleteRepository(deleteuser *User) error
	UserDeletePhotoRepository(deleteuser *User) error
	UserDeleteCommentRepository(deleteuser *User) error
	UserDeleteSocMedRepository(deleteuser *User) error
	GetUserId(loginEmail *User) (int, error)
}

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegister struct {
	RegisterAge      int    `json:"age"`
	RegisterEmail    string `json:"email"`
	RegisterId       int    `json:"id"`
	RegisterUsername string `json:"username"`
}

type UserUpdate struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Token struct {
	JwtToken string `json:"token"`
}
