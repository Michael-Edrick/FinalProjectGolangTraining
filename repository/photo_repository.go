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

func (r photoRepository) PhotoPostRepository(postPhoto entity.Photo) (entity.Photo, error) {
	sqlStatement := `
	INSERT INTO photos (title, caption, photo_url, user_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING photoId, user_id, created_at
	`
	rows, err := r.db.Query(sqlStatement, postPhoto.Title, postPhoto.Caption, postPhoto.Photo_url, postPhoto.User_id, time.Now().Local(), time.Now().Local())
	if err != nil {
		return entity.Photo{}, err
	}
	for rows.Next() {
		err = rows.Scan(&postPhoto.Id, &postPhoto.User_id, &postPhoto.Created_at)
		if err != nil {
			return entity.Photo{}, err
		}
	}
	return postPhoto, nil
}

func (r photoRepository) PhotoGetRepository(getPhotos entity.Photo) ([]entity.PhotoGet, error) {
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
	rows, err := r.db.Query(sqlStatement, getPhotos.User_id)
	if err != nil {
		return []entity.PhotoGet{}, err
	}
	for rows.Next() {
		var photoGet entity.PhotoGet
		err = rows.Scan(&photoGet.Id, &photoGet.Title, &photoGet.Caption, &photoGet.Photo_url, &photoGet.User_id, &photoGet.Created_at, &photoGet.Updated_at, &photoGet.User.Email, &photoGet.User.Username)
		if err != nil {
			return []entity.PhotoGet{}, err
		}
		photos = append(photos, photoGet)
	}
	return photos, nil
}

func (r photoRepository) PhotoUpdateRepository(updatePhoto entity.Photo) (entity.Photo, error) {
	sqlStatement := `
	UPDATE photos 
	SET title = $1, caption = $2, photo_url = $3, updated_at = $4
	WHERE photoId = $5
	RETURNING photoId, title, caption, photo_url, user_id, updated_at
	`
	rows, err := r.db.Query(sqlStatement, updatePhoto.Title, updatePhoto.Caption, updatePhoto.Photo_url, time.Now().Local(), updatePhoto.Id)
	if err != nil {
		return entity.Photo{}, err
	}
	for rows.Next() {
		err = rows.Scan(&updatePhoto.Id, &updatePhoto.Title, &updatePhoto.Caption, &updatePhoto.Photo_url, &updatePhoto.User_id, &updatePhoto.Updated_at)
		if err != nil {
			return entity.Photo{}, err
		}
	}
	return updatePhoto, nil
}

func (r photoRepository) PhotoDeleteRepository(deletePhoto entity.Photo) error {
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
