package mapper

import "FinalProject/entity"

func PostPhotoMapper(response entity.Photo)entity.PhotoPost{
	var photoPost entity.PhotoPost
	photoPost.Id = response.Id
	photoPost.Title = response.Title
	photoPost.Caption = response.Caption
	photoPost.Photo_url = response.Photo_url
	photoPost.User_id = response.User_id
	photoPost.Created_at = response.Created_at
	return photoPost
}

func UpdatePhotoMapper(response entity.Photo)entity.PhotoUpdate{
	var photoUpdate entity.PhotoUpdate
	photoUpdate.Id = response.Id
	photoUpdate.Title = response.Title
	photoUpdate.Caption = response.Caption
	photoUpdate.Photo_url = response.Photo_url
	photoUpdate.User_id = response.User_id
	photoUpdate.Updated_at = response.Updated_at
	return photoUpdate
}