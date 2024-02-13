package db

import "blog/models"

func GetPostsFromDB(page, perPage int) ([]models.Post, int, error) {
	offset := (page - 1) * perPage
	limit := perPage

	var p models.Post
	var posts []models.Post
	// SQL-запрос на получение данных из таблицы posts
	query := `SELECT * FROM posts ORDER BY post_id DESC LIMIT $1 OFFSET $2`
	rows, err := DB.Query(query, limit, offset)
	if err != nil {
		return posts, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&p.PostId, &p.UserId, &p.CategoryId, &p.PostTitle, &p.PostContent, &p.CreatedAt)
		// if err != nil {
		// 	return nil, 0, err
		// }
		posts = append(posts, p)
	}

	// Получение общего количества записей
	countQuery := "SELECT COUNT(*) FROM posts"
	var totalCount int
	err = DB.QueryRow(countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return posts, totalCount, nil
}
