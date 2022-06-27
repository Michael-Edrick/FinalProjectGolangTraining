package entity

import "time"

type PhotoRepositoryInterface interface{

}

type PhotoServiceInterface interface{

}

type Photo struct {
	Id         uint
	Title      string
	Caption    string
	Photo_url  string
	User_id    uint
	Created_at time.Time
	Updated_at time.Time
}