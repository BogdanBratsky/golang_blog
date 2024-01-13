package db

import "database/sql"

func UserExists(uEmail, uName, uPass *string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE user_email = $1 AND user_name = $2 and password_hash = $3`
	var count int
	err := DB.QueryRow(query, uEmail, uName, uPass).Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если нет строк, значит, почтовый адрес не существует
			return false, nil
		}
		// В случае другой ошибки возвращаем ошибку
		return false, err
	}
	return count > 0, nil
}
