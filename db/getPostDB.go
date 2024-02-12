package db

import (
	"database/sql"
	"log"
	"time"
)

func GetPostFromDB(id *int) (Post, error) {
	var p Post
	var createdAt sql.NullTime
	query := `SELECT * FROM posts WHERE post_id = $1`
	err := DB.QueryRow(query, *id).Scan(&p.PostId, &p.UserId, &p.CategoryId, &p.PostTitle, &p.PostContent, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, err
		} else {
			log.Fatal(err)
		}
	}

	// Проверка, является ли createdAt NULL
	if createdAt.Valid {
		p.CreatedAt = createdAt.Time
	} else {
		// Установка значения по умолчанию (или другого значения по вашему выбору)
		p.CreatedAt = time.Time{} // Например, можно использовать time.Now() или другую подходящую дату
	}

	return p, nil
}
