package service

import (
	"FinalProject/entity"
	"FinalProject/utils"
	"errors"
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository entity.UserRepository
}

func NewUserService(userRepository entity.UserRepository) entity.UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s UserService) RegisterService(newUser entity.User) (entity.User, error) {
	//validasi register
	email := newUser.Email
	_,err := mail.ParseAddress(email)
	if newUser.Email == ""{
		return entity.User{}, errors.New("email must be filled")
	}
	if err != nil{
		return entity.User{}, errors.New("email not valid")
	}
	if newUser.Username == ""{
		return entity.User{}, errors.New("username must be filled")
	}
	if newUser.Password == "" || len(newUser.Password) < 6{
		return entity.User{}, errors.New("password must be filled and must be longer than 6 characters")
	}
	if newUser.Age <= 8{
		return entity.User{}, errors.New("age must not below 8")
	}

	//hash password
	newUser.Password = hashPassword(newUser.Password)
	return s.userRepository.UserRegisterRepository(newUser)
}

func (s UserService)LoginService(newLogin entity.User) (string, error){
	pass, err := s.userRepository.UserLoginRepository(newLogin)
	//validasi check email dan password
	if err != nil{
		return "", err
	}
	err = checkPassword(pass, newLogin.Password)
	if err != nil{
		return "", errors.New("password didn't match")
	}
	jwtToken := entity.Token{}
	token, _ := utils.GenerateJWT(newLogin.Email)
	jwtToken.JwtToken = token
	return jwtToken.JwtToken, nil
}

func (s UserService)UserUpdateService(updateUser entity.User)(entity.User, error){
	email := updateUser.Email
	_,err := mail.ParseAddress(email)
	if updateUser.Email == ""{
		return entity.User{}, errors.New("email must be filled")
	}
	if err != nil{
		return entity.User{}, errors.New("email not valid")
	}
	if updateUser.Username == ""{
		return entity.User{}, errors.New("username must be filled")
	}
	return s.userRepository.UserUpdateRepository(updateUser)
}
func (s UserService)UserDeleteService(loginEmail entity.User)error{
	return s.userRepository.UserDeleteRepository(loginEmail)
}

func hashPassword(password string) string{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	fmt.Println(string(hashedPassword))
	return string(hashedPassword)
}

func checkPassword(dbPassword string, password string) error{
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return err
	} 
	return nil
}



