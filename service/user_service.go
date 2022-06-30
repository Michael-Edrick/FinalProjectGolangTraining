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
	userRepository entity.UserRepositoryInterface
}

func NewUserService(userRepository entity.UserRepositoryInterface) entity.UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s UserService) UserRegisterService(newUser *entity.User) (*entity.User, error) {
	//validasi register
	email := newUser.Email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, errors.New("email not valid")
	}
	if newUser.Email == "" {
		return nil, errors.New("email must be filled")
	}
	if newUser.Username == "" {
		return nil, errors.New("username must be filled")
	}
	if newUser.Password == "" || len(newUser.Password) < 6 {
		return nil, errors.New("password must be filled and must be longer than 6 characters")
	}
	if newUser.Age <= 8 {
		return nil, errors.New("age must not below 8")
	}

	//hash password
	newUser.Password = hashPassword(newUser.Password)
	return s.userRepository.UserRegisterRepository(newUser)
}

func (s UserService) UserLoginService(newLogin *entity.User) (string, error) {
	var data *entity.User
	data, err := s.userRepository.UserLoginRepository(newLogin)
	//validasi check email dan password
	if err != nil {
		return "", err
	}
	err = checkPassword(data.Password, newLogin.Password)
	if err != nil {
		return "", errors.New("password didn't match")
	}
	jwtToken := entity.Token{}
	token, _ := utils.GenerateJWT(data.Id)
	jwtToken.JwtToken = token
	return jwtToken.JwtToken, nil
}

func (s UserService) UserUpdateService(updateUser *entity.User) (*entity.User, error) {
	email := updateUser.Email
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, errors.New("email not valid")
	}
	if updateUser.Email == "" {
		return nil, errors.New("email must be filled")
	}
	if updateUser.Username == "" {
		return nil, errors.New("username must be filled")
	}
	return s.userRepository.UserUpdateRepository(updateUser)
}
func (s UserService) UserDeleteService(loginUser *entity.User) error {

	err := s.userRepository.UserDeletePhotoRepository(loginUser)
	if err != nil {
		return errors.New("something went wrong")
	}
	err = s.userRepository.UserDeleteCommentRepository(loginUser)
	if err != nil {
		return errors.New("something went wrong")
	}
	err = s.userRepository.UserDeleteSocMedRepository(loginUser)
	if err != nil {
		return errors.New("something went wrong")
	}
	err = s.userRepository.UserDeleteRepository(loginUser)
	if err != nil {
		return errors.New("something went wrong")
	}
	return nil
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	fmt.Println(string(hashedPassword))
	return string(hashedPassword)
}

func checkPassword(dbPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
