package service

import (
	"FinalProject/entity"
	"errors"
)

type CommentService struct {
	commentRepository entity.CommentRepositoryInterface
}

func NewCommentService(commentRepository entity.CommentRepositoryInterface) entity.CommentServiceInterface {
	return &CommentService{
		commentRepository: commentRepository,
	}
}

func (s CommentService) CommentPostService(postComment *entity.Comment) (*entity.Comment, error) {
	if postComment.Message == "" {
		return nil, errors.New("message must be filled")
	}
	return s.commentRepository.CommentPostRepository(postComment)
}

func (s CommentService) CommentGetService(getComment *entity.Comment) ([]entity.CommentGet, error) {
	return s.commentRepository.CommentGetRepository(getComment)
}

func (s CommentService) CommentUpdateService(updateComment *entity.Comment) (*entity.CommentUpdate, error) {
	if updateComment.Message == "" {
		return nil, errors.New("message must be filled")
	}
	return s.commentRepository.CommentUpdateRepository(updateComment)
}

func (s CommentService) CommentDeleteService(deleteComment *entity.Comment) error {
	err := s.commentRepository.CommentDeleteRepository(deleteComment)
	if err != nil {
		return errors.New("something went wrong")
	}
	return nil
}
