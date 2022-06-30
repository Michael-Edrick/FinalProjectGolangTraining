package entity

import "time"

type SocialMediaServiceInterface interface {
	SocialMediaPostService(postSocialMedia SocialMedia) (SocialMedia, error)
	SocialMediaGetService(getSocialMedia SocialMedia) ([]SocialMediaGetData, error)
	SocialMediaUpdateService(updateSocialMedia SocialMedia) (SocialMedia, error)
	SocialMediaDeleteService(deleteSocialMedia SocialMedia) error
}

type SocialMediaRepositoryInterface interface {
	SocialMediaPostRepository(postSocialMedia SocialMedia) (SocialMedia, error)
	SocialMediaGetRepository(getSocialMedia SocialMedia) ([]SocialMediaGetData, error)
	SocialMediaUpdateRepository(updateSocialMedia SocialMedia) (SocialMedia, error)
	SocialMediaDeleteRepository(deleteSocialMedia SocialMedia) error
}

type SocialMedia struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	Created_at       time.Time `json:"created_at"`
	Updated_at       time.Time `json:"updated_at`
}

type SocialMediaPost struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	Created_at       time.Time `json:"created_at"`
}

type SocialMediaGet struct {
	SocialMedias []SocialMediaGetData `json:"social_medias"`
}

type SocialMediaGetData struct {
	Id               int                `json:"id"`
	Name             string             `json:"name"`
	Social_media_url string             `json:"social_media_url"`
	User_id          int                `json:"UserId"`
	Created_at       time.Time          `json:"createdAt"`
	Updated_at       time.Time          `json:"updatedAt"`
	User             UserSocialMediaGet `json:"User"`
}

type UserSocialMediaGet struct {
	Id                int    `json:"id"`
	Username          string `json:"username"`
	Profile_image_url string `json:"profile_image_url"`
}

type SocialMediaUpdate struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Social_media_url string    `json:"social_media_url"`
	User_id          int       `json:"user_id"`
	Updated_at       time.Time `json:"updated_at"`
}
