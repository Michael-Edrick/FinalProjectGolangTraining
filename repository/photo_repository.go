package repository

import (
	"FinalProject/entity"
	"database/sql"
)

type photoRepository struct {
	db *sql.DB
}

func NewPhotoRepository(db *sql.DB) entity.PhotoRepositoryInterface {
	return photoRepository{
		db: db,
	}
}