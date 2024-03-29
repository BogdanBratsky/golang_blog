package db

import "blog/models"

func GetArticlesByCategoryDB(id, page, perPage *int) ([]models.Post, int, error) {
	offset := (*page - 1) * *perPage
	limit := *perPage

	var p models.Post
	var posts []models.Post
	query := `SELECT * FROM posts WHERE category_id = $1 ORDER BY post_id DESC LIMIT $2 OFFSET $3`
	rows, err := DB.Query(query, *id, limit, offset)
	if err != nil {
		return posts, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.PostId, &p.UserId, &p.CategoryId, &p.PostTitle, &p.PostContent, &p.CreatedAt)
		posts = append(posts, p)
	}

	// Получение общего количества записей
	countQuery := "SELECT COUNT(*) FROM posts WHERE category_id = $1"
	var totalCount int
	err = DB.QueryRow(countQuery, *id).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return posts, totalCount, nil
}
