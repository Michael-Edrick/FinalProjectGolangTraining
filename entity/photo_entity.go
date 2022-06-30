package entity

import "time"

type PhotoServiceInterface interface {
	PhotoPostService(postPhoto *Photo) (*Photo, error)
	PhotoGetService(getPhotos *Photo) ([]PhotoGet, error)
	PhotoUpdateService(updatePhoto *Photo) (*Photo, error)
	PhotoDeleteService(deletePhoto *Photo) error
}

type PhotoRepositoryInterface interface {
	PhotoPostRepository(postPhoto *Photo) (*Photo, error)
	PhotoGetRepository(getPhotos *Photo) ([]PhotoGet, error)
	PhotoUpdateRepository(updatePhoto *Photo) (*Photo, error)
	PhotoDeleteRepository(deletePhoto *Photo) error
}

type Photo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at`
}

type PhotoPost struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoGet struct {
	Id        int          `json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserId    int          `json:"user_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at`
	User      UserPhotoGet `json:"User"`
}

type UserPhotoGet struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoUpdate struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
