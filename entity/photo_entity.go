package entity

import "time"

type PhotoServiceInterface interface{
	PhotoPostService(postPhoto Photo, loginUser User)(Photo, error)
	PhotoGetService(loginUser User)([]PhotoGet, error)
	PhotoUpdateService(updatePhoto Photo)(Photo, error)
	PhotoDeleteService(deletePhoto Photo)(error)
}

type PhotoRepositoryInterface interface{
	PhotoPostRepository(postPhoto Photo, loginUser User)(Photo, error)
	PhotoGetRepository(loginUser User)([]PhotoGet, error)
	PhotoUpdateRepository(updatePhoto Photo)(Photo, error)
	PhotoDeleteRepository(deletePhoto Photo)(error)
}

type Photo struct {
	Id         int `json:"id"`
	Title      string `json:"title"`
	Caption    string	`json:"caption"`
	Photo_url  string	`json:"photo_url"`
	User_id    int	`json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at`
}

type PhotoPost struct {
	Id         int `json:"id"`
	Title      string `json:"title"`
	Caption    string	`json:"caption"`
	Photo_url  string	`json:"photo_url"`
	User_id    int	`json:"user_id"`
	Created_at time.Time `json:"created_at"`
}

type PhotoGet struct {
	Id         int `json:"id"`
	Title      string `json:"title"`
	Caption    string	`json:"caption"`
	Photo_url  string	`json:"photo_url"`
	User_id    int	`json:"user_id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at`
	User UserPhotoGet `json:"User"`
}

type UserPhotoGet struct {
	Email    string	`json:"email"`
	Username string `json:"username"`
}

type PhotoUpdate struct {
	Id         int `json:"id"`
	Title      string `json:"title"`
	Caption    string	`json:"caption"`
	Photo_url  string	`json:"photo_url"`
	User_id    int	`json:"user_id"`
	Updated_at time.Time `json:"updated_at"`
}