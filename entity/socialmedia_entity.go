package entity

import "time"

type SocialMediaServiceInterface interface {
	SocialMediaPostService(postSocialMedia *SocialMedia) (*SocialMedia, error)
	SocialMediaGetService(getSocialMedia *SocialMedia) ([]SocialMediaGetData, error)
	SocialMediaUpdateService(updateSocialMedia *SocialMedia) (*SocialMedia, error)
	SocialMediaDeleteService(deleteSocialMedia *SocialMedia) error
}

type SocialMediaRepositoryInterface interface {
	SocialMediaPostRepository(postSocialMedia *SocialMedia) (*SocialMedia, error)
	SocialMediaGetRepository(getSocialMedia *SocialMedia) ([]SocialMediaGetData, error)
	SocialMediaUpdateRepository(updateSocialMedia *SocialMedia) (*SocialMedia, error)
	SocialMediaDeleteRepository(deleteSocialMedia *SocialMedia) error
}

type SocialMedia struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at`
}

type SocialMediaPost struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaGet struct {
	SocialMedias []SocialMediaGetData `json:"social_medias"`
}

type SocialMediaGetData struct {
	Id             int                `json:"id"`
	Name           string             `json:"name"`
	SocialMediaUrl string             `json:"social_media_url"`
	UserId         int                `json:"UserId"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	User           UserSocialMediaGet `json:"User"`
}

type UserSocialMediaGet struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type SocialMediaUpdate struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"user_id"`
	UpdatedAt      time.Time `json:"updated_at"`
}
