package db

func GetPostsFromDB() ([]Post, error) {
	// sql-запрос на получение данных из таблицы posts
	rows, err := DB.Query("SELECT * FROM posts ORDER BY post_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var p Post
		err := rows.Scan(&p.PostId, &p.UserId, &p.CategoryId, &p.PostTitle, &p.PostContent, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}
