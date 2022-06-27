package repository

import (
	"FinalProject/entity"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type userRepository struct{
	db *sql.DB
}

func NewUserRepository (db *sql.DB) entity.UserRepositoryInterface{
	return userRepository{
		db:db,
	}
}

func (r userRepository)UserRegisterRepository(newUser entity.User) (entity.User, error) {
	sqlEmailUnique :=`
	SELECT userId FROM users WHERE email = $1
	`
	rows, err := r.db.Query(sqlEmailUnique, newUser.Email)
    if err != nil {
		return entity.User{}, err
	}
	if rows.Next() {
        return entity.User{}, errors.New("email already registered")
    }

	sqlUsernameUnique :=`
	SELECT userId FROM users WHERE username = $1
	`
	rows, err = r.db.Query(sqlUsernameUnique, newUser.Username)
    if err != nil {
		return entity.User{}, err
	}
	if rows.Next() {
        return entity.User{}, errors.New("username already exists")
    }
	
	sqlStatement := `
	INSERT INTO users (username, email, password, age, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING userId
	`
	rows, err = r.db.Query(sqlStatement, newUser.Username, newUser.Email, newUser.Password, newUser.Age, time.Now().Local(), time.Now().Local())
	if err != nil {
		return entity.User{}, err
	}
	for rows.Next() {
		err = rows.Scan(&newUser.Id)
		if err != nil {
			return entity.User{}, err
		}
	}

	return newUser, nil 
}

func (r userRepository)UserLoginRepository(newLogin entity.User) (string, error) {
	var data entity.User
	sqlStatement := `
	SELECT password FROM users 
	WHERE email = $1
	`
	rows, err := r.db.Query(sqlStatement, newLogin.Email)
	if err != nil {
		return "", err
	}
	//validate email exists by rows
	for rows.Next() {
		err = rows.Scan(&data.Password)

		if err!= nil{
			return "", err
		}
	}
	if data.Password == "" {
		return "", errors.New("email doesn't registered")
	}
	return data.Password, nil
}

func (r userRepository)UserUpdateRepository(updateUser entity.User)(entity.User, error){
	sqlStatement := `
	UPDATE users
	SET username = $1 , email = $2, updated_at = $3
	WHERE userId = $4
	RETURNING userId, email, username, age, updated_at
	`
	fmt.Printf("%+v\n",updateUser)
	rows, err := r.db.Query(sqlStatement, updateUser.Username, updateUser.Email, time.Now().Local(), updateUser.Id)
	if err != nil {
		return entity.User{}, err
	}
	for rows.Next() {
		err = rows.Scan(&updateUser.Id, &updateUser.Email, &updateUser.Username, &updateUser.Age, &updateUser.Updated_at)
		if err != nil {
			return entity.User{}, err
		}
	}

	return updateUser, nil
}

func (r userRepository)UserDeleteRepository(loginEmail entity.User)error{
	sqlStatement :=`
	DELETE FROM users
	WHERE userId = $1
	`

	_, err := r.db.Exec(sqlStatement, loginEmail.Id)
	if err != nil {
		return err
	}
	
	return err
}

func (r userRepository)UserDeletePhotoRepository(loginEmail entity.User)error{
	sqlStatement :=`
	DELETE FROM photos
	WHERE user_id = $1
	`

	_, err := r.db.Exec(sqlStatement, loginEmail.Id)
	if err != nil {
		return err
	}
	
	return err
}
func (r userRepository)UserDeleteCommentRepository(loginEmail entity.User)error{
	sqlStatement :=`
	DELETE FROM comments
	WHERE user_id = $1
	`

	_, err := r.db.Exec(sqlStatement, loginEmail.Id)
	if err != nil {
		return err
	}
	
	return err
}
func (r userRepository)UserDeleteSocMedRepository(loginEmail entity.User)error{
	sqlStatement :=`
	DELETE FROM socialmedia
	WHERE user_id = $1
	`

	_, err := r.db.Exec(sqlStatement, loginEmail.Id)
	if err != nil {
		return err
	}
	
	return err
}

func (r userRepository)GetUserId(loginEmail entity.User)(int ,error){
	sqlStatement := `
	SELECT userID
	FROM users
	WHERE email = $1
	`

	rows, err := r.db.Query(sqlStatement, loginEmail.Email)
	if err!= nil {
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&loginEmail.Id)
		if err != nil {
			return 0, err
		}
	}
	return loginEmail.Id, nil
}