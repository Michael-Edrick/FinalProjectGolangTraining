package service

import (
	"FinalProject/entity"
	"errors"
)

type PhotoService struct {
	photoRepository entity.PhotoRepositoryInterface
}

func NewPhotoService(photoRepository entity.PhotoRepositoryInterface) entity.PhotoServiceInterface {
	return &PhotoService{
		photoRepository: photoRepository,
	}
}

func (s PhotoService) PhotoPostService(postPhoto *entity.Photo) (*entity.Photo, error) {
	if postPhoto.Title == "" {
		return nil, errors.New("photo title must be filled")
	}
	if postPhoto.PhotoUrl == "" {
		return nil, errors.New("photo url must be filled")
	}
	return s.photoRepository.PhotoPostRepository(postPhoto)
}

func (s PhotoService) PhotoGetService(getPhotos *entity.Photo) ([]entity.PhotoGet, error) {
	return s.photoRepository.PhotoGetRepository(getPhotos)
}

func (s PhotoService) PhotoUpdateService(updatePhoto *entity.Photo) (*entity.Photo, error) {
	if updatePhoto.Title == "" {
		return nil, errors.New("photo title must be filled")
	}
	if updatePhoto.PhotoUrl == "" {
		return nil, errors.New("photo url must be filled")
	}
	return s.photoRepository.PhotoUpdateRepository(updatePhoto)
}

func (s PhotoService) PhotoDeleteService(deletePhoto *entity.Photo) error {
	err := s.photoRepository.PhotoDeleteRepository(deletePhoto)
	if err != nil {
		return errors.New("something went wrong")
	}
	return nil
}
