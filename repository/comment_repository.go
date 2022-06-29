package repository

import (
	"FinalProject/entity"
	"database/sql"
	"time"
)

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) entity.CommentRepositoryInterface {
	return commentRepository{
		db: db,
	}
}

func (r commentRepository) CommentPostRepository(postComment entity.Comment) (entity.Comment, error) {
	sqlStatement := `
	INSERT INTO comments (message, photo_id, user_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING commentId, created_at
	`
	rows, err := r.db.Query(sqlStatement, postComment.Message, postComment.Photo_id, postComment.User_id, time.Now().Local(), time.Now().Local())
	if err != nil {
		return entity.Comment{}, err
	}
	for rows.Next() {
		err = rows.Scan(&postComment.Id, &postComment.Created_at)
		if err != nil {
			return entity.Comment{}, err
		}
	}
	return postComment, nil
}

func (r commentRepository) CommentGetRepository(getComment entity.Comment) ([]entity.CommentGet, error) {
	var response []entity.CommentGet
	sqlStatement := `
	SELECT 
		c.commentId,
		c.message,
		c.photo_id,
		c.user_id,
		c.updated_at,
		c.created_at,
		u.userId,
		u.email,
		u.username,
		p.photoId,
		p.title,
		p.caption, 
		p.photo_url,
		p.user_id
	FROM comments as c 
	INNER JOIN users as u on (c.user_id = u.userId)
	INNER JOIN photos as p on (c.photo_id = p.photoId)
	WHERE u.userId = $1
	`
	rows, err := r.db.Query(sqlStatement, getComment.User_id)
	if err != nil {
		return []entity.CommentGet{}, err
	}
	for rows.Next() {
		var commentGet entity.CommentGet
		err = rows.Scan(&commentGet.Id, &commentGet.Message, &commentGet.Photo_id, &commentGet.User_id, &commentGet.Updated_at, &commentGet.Created_at, &commentGet.User.Id, &commentGet.User.Email, &commentGet.User.Username, &commentGet.Photo.Id, &commentGet.Photo.Title, &commentGet.Photo.Caption, &commentGet.Photo.Photo_url, &commentGet.Photo.User_id)
		if err != nil {
			return []entity.CommentGet{}, err
		}
		response = append(response, commentGet)
	}
	return response, nil
}

func (r commentRepository) CommentUpdateRepository(updateComment entity.Comment) (entity.CommentUpdate, error) {
	var response entity.CommentUpdate
	sqlStatement := `
	UPDATE comments as c
	SET message = $1, updated_at = $2
	FROM photos as p
	WHERE c.commentId = $3 AND p.photoId = c.photo_id
	RETURNING c.commentId, p.title, p.caption, p.photo_url, c.user_id, c.updated_at
	`
	rows, err := r.db.Query(sqlStatement, updateComment.Message, time.Now().Local(), updateComment.Id)
	if err != nil {
		return entity.CommentUpdate{}, err
	}
	for rows.Next() {
		err = rows.Scan(&response.Id, &response.Title, &response.Caption, &response.Photo_url, &response.User_id, &response.Updated_at)
		if err != nil {
			return entity.CommentUpdate{}, err
		}
	}
	return response, nil
}

func (r commentRepository) CommentDeleteRepository(deleteComment entity.Comment) error {
	sqlStatement := `
	DELETE FROM comments
	WHERE commentId = $1
	`
	_, err := r.db.Exec(sqlStatement, deleteComment.Id)
	if err != nil {
		return err
	}
	return err
}
