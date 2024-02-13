package db

import "blog/models"

func CreatePostDB(p *models.Post) error {
	var err error
	// SQL-запрос для вставки пользователя
	query := `INSERT INTO posts (post_title, post_content, user_id, category_id, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING post_id`
	_, err = DB.Exec(query, p.PostTitle, p.PostContent, p.UserId, p.CategoryId, p.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
