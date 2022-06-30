package service

import (
	"FinalProject/entity"
	"errors"
)

type SocialMediaService struct {
	socialMediaRepository entity.SocialMediaRepositoryInterface
}

func NewSocialMediaService(socialMediaRepository entity.SocialMediaRepositoryInterface) entity.SocialMediaServiceInterface {
	return &SocialMediaService{
		socialMediaRepository: socialMediaRepository,
	}
}

func (s SocialMediaService) SocialMediaPostService(postSocialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	if postSocialMedia.Name == "" {
		return nil, errors.New("name must be filled")
	}
	if postSocialMedia.SocialMediaUrl == "" {
		return nil, errors.New("social media url must be filled")
	}
	return s.socialMediaRepository.SocialMediaPostRepository(postSocialMedia)
}

func (s SocialMediaService) SocialMediaGetService(getSocialMedia *entity.SocialMedia) ([]entity.SocialMediaGetData, error) {
	return s.socialMediaRepository.SocialMediaGetRepository(getSocialMedia)
}

func (s SocialMediaService) SocialMediaUpdateService(updateSocialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	if updateSocialMedia.Name == "" {
		return nil, errors.New("name must be filled")
	}
	if updateSocialMedia.SocialMediaUrl == "" {
		return nil, errors.New("social media url must be filled")
	}
	return s.socialMediaRepository.SocialMediaUpdateRepository(updateSocialMedia)
}

func (s SocialMediaService) SocialMediaDeleteService(deleteSocialMedia *entity.SocialMedia) error {
	err := s.socialMediaRepository.SocialMediaDeleteRepository(deleteSocialMedia)
	if err != nil {
		return errors.New("something went wrong")
	}
	return nil
}
