package db

func CreatePostDB(p *Post) error {
	var err error
	// SQL-запрос для вставки пользователя
	_, err = DB.Exec(`INSERT INTO posts (post_title, post_content) VALUES ($1, $2) RETURNING post_id`,
		p.PostTitle, p.PostContent)
	if err != nil {
		return err
	}
	return nil
}
