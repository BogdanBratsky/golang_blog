package db

import (
	"database/sql"
	"fmt"
)

func UserExists(uEmail, uPass *string) (uint64, string, error) {
	query := `SELECT user_id, user_name FROM users WHERE user_email = $1 AND password_hash = $2`
	var usId uint64
	var usName string
	err := DB.QueryRow(query, uEmail, uPass).Scan(&usId, &usName)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если нет строк, значит, почтовый адрес не существует
			return 0, "", nil
		}
		// В случае другой ошибки возвращаем ошибку
		return 0, "", err
	}
	fmt.Println(usId)
	return usId, usName, nil
}
