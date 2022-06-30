package repository

import (
	"FinalProject/entity"
	"database/sql"
	"time"
)

type socialMediaRepository struct {
	db *sql.DB
}

func NewSocialMediaRepository(db *sql.DB) entity.SocialMediaRepositoryInterface {
	return socialMediaRepository{
		db: db,
	}
}

func (r socialMediaRepository) SocialMediaPostRepository(postSocialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	sqlStatement := `
	INSERT INTO socialmedia (name, social_media_url, user_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING scId, created_at
	`
	rows, err := r.db.Query(sqlStatement, postSocialMedia.Name, postSocialMedia.SocialMediaUrl, postSocialMedia.UserId, time.Now().Local(), time.Now().Local())
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&postSocialMedia.Id, &postSocialMedia.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return postSocialMedia, nil
}

func (r socialMediaRepository) SocialMediaGetRepository(getSocialMedia *entity.SocialMedia) ([]entity.SocialMediaGetData, error) {
	var response []entity.SocialMediaGetData
	sqlStatement := `
		SELECT 
		s.scId,
		s.name,
		s.social_media_url, 
		s.user_id,
		s.created_at,
		s.updated_at,
		u.userId,
		u.username,
		p.title
	FROM socialmedia as s
	INNER JOIN users as u ON (s.user_id = u.userId)
	INNER JOIN photos as p ON (u.userId = p.user_id)
	where u.userId = $1 AND p.title = 'profile-image.com'
	`
	rows, err := r.db.Query(sqlStatement, getSocialMedia.UserId)
	if err != nil {
		return []entity.SocialMediaGetData{}, err
	}
	for rows.Next() {
		var socialMediaGetData entity.SocialMediaGetData
		err = rows.Scan(&socialMediaGetData.Id, &socialMediaGetData.Name, &socialMediaGetData.SocialMediaUrl, &socialMediaGetData.UserId, &socialMediaGetData.CreatedAt, &socialMediaGetData.UpdatedAt, &socialMediaGetData.User.Id, &socialMediaGetData.User.Username, &socialMediaGetData.User.ProfileImageUrl)
		if err != nil {
			return []entity.SocialMediaGetData{}, err
		}
		response = append(response, socialMediaGetData)
	}
	return response, nil
}

func (r socialMediaRepository) SocialMediaUpdateRepository(updateSocialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	sqlStatement := `
	UPDATE socialmedia 
	SET name = $1, social_media_url = $2, updated_at = $3
	WHERE scId = $4
	RETURNING user_id, updated_at
	`
	rows, err := r.db.Query(sqlStatement, updateSocialMedia.Name, updateSocialMedia.SocialMediaUrl, time.Now().Local(), updateSocialMedia.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&updateSocialMedia.UserId, &updateSocialMedia.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return updateSocialMedia, nil
}

func (r socialMediaRepository) SocialMediaDeleteRepository(deleteSocialMedia *entity.SocialMedia) error {
	sqlStatement := `
	DELETE FROM socialmedia
	WHERE scId = $1
	`
	_, err := r.db.Exec(sqlStatement, deleteSocialMedia.Id)
	if err != nil {
		return err
	}
	return err
}
