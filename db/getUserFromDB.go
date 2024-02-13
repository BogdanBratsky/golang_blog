package db

import (
	"blog/models"
	"database/sql"
	"log"
)

func GetUserFromDB(id *int) (models.Users, error) {
	var u models.Users
	query := `SELECT user_id, user_name, user_email, created_at FROM users WHERE user_id = $1`
	err := DB.QueryRow(query, *id).Scan(&u.UserId, &u.UserName, &u.UserEmail, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, err
		}
		log.Println("Error executing query:", err)
		return u, err
	}
	return u, nil
}
