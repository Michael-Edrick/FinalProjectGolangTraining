package entity

type SocialMediaServiceInterface interface {
}

type SocialMediaRepositoryInterface interface {
}

type SocialMedia struct {
	Id               uint
	Name             string
	Social_media_url string
	User_id          uint
}
