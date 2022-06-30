package mapper

import "FinalProject/entity"

func PostPhotoMapper(response *entity.Photo) *entity.PhotoPost {
	var photoPost entity.PhotoPost
	photoPost.Id = response.Id
	photoPost.Title = response.Title
	photoPost.Caption = response.Caption
	photoPost.PhotoUrl = response.PhotoUrl
	photoPost.UserId = response.UserId
	photoPost.CreatedAt = response.CreatedAt
	return &photoPost
}

func UpdatePhotoMapper(response *entity.Photo) *entity.PhotoUpdate {
	var photoUpdate entity.PhotoUpdate
	photoUpdate.Id = response.Id
	photoUpdate.Title = response.Title
	photoUpdate.Caption = response.Caption
	photoUpdate.PhotoUrl = response.PhotoUrl
	photoUpdate.UserId = response.UserId
	photoUpdate.UpdatedAt = response.UpdatedAt
	return &photoUpdate
}
