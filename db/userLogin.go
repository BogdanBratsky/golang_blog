package db

import (
	"database/sql"
	"fmt"
)

func UserExists(uEmail, uPass *string) (uint64, error) {
	query := `SELECT user_id FROM users WHERE user_email = $1 AND password_hash = $2`
	var usId uint64
	err := DB.QueryRow(query, uEmail, uPass).Scan(&usId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если нет строк, значит, почтовый адрес не существует
			return 0, nil
		}
		// В случае другой ошибки возвращаем ошибку
		return 0, err
	}
	fmt.Println(usId)
	return usId, nil
}
