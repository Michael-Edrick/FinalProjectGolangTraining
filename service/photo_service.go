package service

import "FinalProject/entity"

type PhotoService struct {
	photoRepository entity.PhotoRepositoryInterface
}

func NewPhotoService(photoRepository entity.PhotoRepositoryInterface) entity.PhotoServiceInterface {
	return &PhotoService{
		photoRepository: photoRepository,
	}
}
