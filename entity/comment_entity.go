package entity

import "time"

type Comment struct {
	Id         uint
	User_id    uint
	Photo_id   uint
	Message    string
	Created_at time.Time
	Updated_at time.Time
}