package repository

import (
	"FinalProject/entity"
	"database/sql"
	"errors"
	"fmt"
)

type userRepository struct{
	db *sql.DB
}

func NewUserRepository (db *sql.DB) entity.UserRepository{
	return userRepository{
		db:db,
	}
}

func (r userRepository)UserRegisterRepository(newUser entity.User) (entity.User, error) {
	sqlStatement := `
	INSERT INTO users (username, email, password, age, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING userId
	`
	rows, err := r.db.Query(sqlStatement, newUser.Username, newUser.Email, newUser.Password, newUser.Age, newUser.Created_at, newUser.Updated_at)
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
	SET username = $1 , email = $2
	WHERE userId = $3
	RETURNING userId, email, username, age, updated_at
	`
	fmt.Printf("%+v\n",updateUser)
	rows, err := r.db.Query(sqlStatement, updateUser.Username, updateUser.Email, updateUser.Id)
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
	WHERE email = $1
	`

	_, err := r.db.Exec(sqlStatement, loginEmail.Email)
	if err != nil {
		return err
	}
	
	return err
}