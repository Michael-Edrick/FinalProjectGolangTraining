package service

import "FinalProject/entity"

type SocialMediaService struct {
	socialMediaRepository entity.SocialMediaRepositoryInterface
}

func NewSocialMediaService(socialMediaRepository entity.SocialMediaRepositoryInterface) entity.SocialMediaServiceInterface {
	return &SocialMediaService{
		socialMediaRepository: socialMediaRepository,
	}
}
