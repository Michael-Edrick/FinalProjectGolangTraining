package repository

import (
	"FinalProject/entity"
	"database/sql"
)

type socialMediaRepository struct {
	db *sql.DB
}

func NewSocialMediaRepository(db *sql.DB) entity.SocialMediaRepositoryInterface {
	return socialMediaRepository{
		db: db,
	}
}
