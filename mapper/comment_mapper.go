package mapper

import "FinalProject/entity"

func PostCommentMapper(response *entity.Comment) *entity.CommentPost {
	var commentPost entity.CommentPost
	commentPost.Id = response.Id
	commentPost.Message = response.Message
	commentPost.PhotoId = response.PhotoId
	commentPost.UserId = response.UserId
	commentPost.CreatedAt = response.CreatedAt
	return &commentPost
}
