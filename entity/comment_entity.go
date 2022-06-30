package entity

import (
	"time"
)

type CommentServiceInterface interface {
	CommentPostService(postComment *Comment) (*Comment, error)
	CommentGetService(getComment *Comment) ([]CommentGet, error)
	CommentUpdateService(updateComment *Comment) (*CommentUpdate, error)
	CommentDeleteService(deleteComment *Comment) error
}

type CommentRepositoryInterface interface {
	CommentPostRepository(postComment *Comment) (*Comment, error)
	CommentGetRepository(getComment *Comment) ([]CommentGet, error)
	CommentUpdateRepository(updateComment *Comment) (*CommentUpdate, error)
	CommentDeleteRepository(deleteComment *Comment) error
}

type Comment struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	PhotoId   int       `json:"photo_Id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentPost struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_Id"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentGet struct {
	Id        int             `json:"id"`
	Message   string          `json:"message"`
	PhotoId   int             `json:"photo_Id"`
	UserId    int             `json:"user_id"`
	UpdatedAt time.Time       `json:"updated_at"`
	CreatedAt time.Time       `json:"created_at"`
	User      UserCommentGet  `json:"User"`
	Photo     PhotoCommentGet `json:"Photo"`
}

type UserCommentGet struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoCommentGet struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

type CommentUpdate struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
