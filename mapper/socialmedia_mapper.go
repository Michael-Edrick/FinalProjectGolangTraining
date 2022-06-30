package mapper

import "FinalProject/entity"

func PostSocialMediaMapper(response *entity.SocialMedia) *entity.SocialMediaPost {
	var socialMediaPost entity.SocialMediaPost
	socialMediaPost.Id = response.Id
	socialMediaPost.Name = response.Name
	socialMediaPost.SocialMediaUrl = response.SocialMediaUrl
	socialMediaPost.UserId = response.UserId
	socialMediaPost.CreatedAt = response.CreatedAt
	return &socialMediaPost
}

func GetSocialMediaMapper(response []entity.SocialMediaGetData) *entity.SocialMediaGet {
	var socialMediaGet entity.SocialMediaGet
	socialMediaGet.SocialMedias = response
	return &socialMediaGet
}

func UpdateSocialMediaMapper(response *entity.SocialMedia) *entity.SocialMediaUpdate {
	var socialMediaUpdate entity.SocialMediaUpdate
	socialMediaUpdate.Id = response.Id
	socialMediaUpdate.Name = response.Name
	socialMediaUpdate.SocialMediaUrl = response.SocialMediaUrl
	socialMediaUpdate.UserId = response.UserId
	socialMediaUpdate.UpdatedAt = response.UpdatedAt
	return &socialMediaUpdate
}
