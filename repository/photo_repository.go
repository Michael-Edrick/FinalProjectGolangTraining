package repository

import (
	"FinalProject/entity"
	"database/sql"
	"time"
)

type photoRepository struct {
	db *sql.DB
}

func NewPhotoRepository(db *sql.DB) entity.PhotoRepositoryInterface {
	return photoRepository{
		db: db,
	}
}

func (r photoRepository) PhotoPostRepository(postPhoto *entity.Photo) (*entity.Photo, error) {
	sqlStatement := `
	INSERT INTO photos (title, caption, photo_url, user_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING photoId, created_at
	`
	rows, err := r.db.Query(sqlStatement, postPhoto.Title, postPhoto.Caption, postPhoto.PhotoUrl, postPhoto.UserId, time.Now().Local(), time.Now().Local())
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&postPhoto.Id, &postPhoto.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return postPhoto, nil
}

func (r photoRepository) PhotoGetRepository(getPhotos *entity.Photo) ([]entity.PhotoGet, error) {
	var photos []entity.PhotoGet
	sqlStatement := `
	SELECT 
		p.photoId,
		p.title,
		p.caption, 
		p.photo_url,
		p.user_id,
		p.created_at,
		p.updated_at,
		u.email,
		u.username
	FROM photos as p 
	LEFT JOIN users as u 
	ON (p.user_id = u.userId)
	WHERE u.userId = $1
	`
	rows, err := r.db.Query(sqlStatement, getPhotos.UserId)
	if err != nil {
		return []entity.PhotoGet{}, err
	}
	for rows.Next() {
		var photoGet entity.PhotoGet
		err = rows.Scan(&photoGet.Id, &photoGet.Title, &photoGet.Caption, &photoGet.PhotoUrl, &photoGet.UserId, &photoGet.CreatedAt, &photoGet.UpdatedAt, &photoGet.User.Email, &photoGet.User.Username)
		if err != nil {
			return []entity.PhotoGet{}, err
		}
		photos = append(photos, photoGet)
	}
	return photos, nil
}

func (r photoRepository) PhotoUpdateRepository(updatePhoto *entity.Photo) (*entity.Photo, error) {
	sqlStatement := `
	UPDATE photos 
	SET title = $1, caption = $2, photo_url = $3, updated_at = $4
	WHERE photoId = $5
	RETURNING user_id, updated_at
	`
	rows, err := r.db.Query(sqlStatement, updatePhoto.Title, updatePhoto.Caption, updatePhoto.PhotoUrl, time.Now().Local(), updatePhoto.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&updatePhoto.UserId, &updatePhoto.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return updatePhoto, nil
}

func (r photoRepository) PhotoDeleteRepository(deletePhoto *entity.Photo) error {
	sqlStatement := `
	DELETE FROM photos
	WHERE photoId = $1
	`
	_, err := r.db.Exec(sqlStatement, deletePhoto.Id)
	if err != nil {
		return err
	}
	return err
}
