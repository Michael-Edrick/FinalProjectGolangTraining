package mapper

import "FinalProject/entity"

func PostSocialMediaMapper(response entity.SocialMedia) entity.SocialMediaPost {
	var socialMediaPost entity.SocialMediaPost
	socialMediaPost.Id = response.Id
	socialMediaPost.Name = response.Name
	socialMediaPost.Social_media_url = response.Social_media_url
	socialMediaPost.User_id = response.User_id
	socialMediaPost.Created_at = response.Created_at
	return socialMediaPost
}

func GetSocialMediaMapper(response []entity.SocialMediaGetData) entity.SocialMediaGet {
	var socialMediaGet entity.SocialMediaGet
	socialMediaGet.SocialMedias = response
	return socialMediaGet
}

func UpdateSocialMediaMapper(response entity.SocialMedia) entity.SocialMediaUpdate {
	var socialMediaUpdate entity.SocialMediaUpdate
	socialMediaUpdate.Id = response.Id
	socialMediaUpdate.Name = response.Name
	socialMediaUpdate.Social_media_url = response.Social_media_url
	socialMediaUpdate.User_id = response.User_id
	socialMediaUpdate.Updated_at = response.Updated_at
	return socialMediaUpdate
}
