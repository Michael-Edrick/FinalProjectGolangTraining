package entity

import "time"

type CommentServiceInterface interface{
	
}

type CommentRepositoryInterface interface{

}

type Comment struct {
	Id         uint
	User_id    uint
	Photo_id   uint
	Message    string
	Created_at time.Time
	Updated_at time.Time
}