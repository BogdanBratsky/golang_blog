package db

import "database/sql"

func CreateUserDB(u *User) error {
	// SQL-запрос для вставки пользователя
	query := `INSERT INTO users (user_name, user_email, password_hash) VALUES ($1, $2, $3) RETURNING user_id`
	_, err := DB.Exec(query, u.UserName, u.UserEmail, u.PasswordHash)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func EmailExists(email *string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE user_email = $1`
	var count int
	err := DB.QueryRow(query, *email).Scan(&count)
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
