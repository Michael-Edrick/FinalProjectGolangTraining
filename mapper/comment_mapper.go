package mapper

import "FinalProject/entity"

func PostCommentMapper(response entity.Comment) entity.CommentPost {
	var commentPost entity.CommentPost
	commentPost.Id = response.Id
	commentPost.Message = response.Message
	commentPost.Photo_id = response.Photo_id
	commentPost.User_id = response.User_id
	commentPost.Created_at = response.Created_at
	return commentPost
}
