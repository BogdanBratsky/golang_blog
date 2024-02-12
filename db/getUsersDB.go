package db

func GetUsersFromDB() ([]Users, error) {
	// sql-запрос на получение данных из таблицы users
	rows, err := DB.Query("SELECT user_id, user_name, user_email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Users
	for rows.Next() {
		var u Users
		err := rows.Scan(&u.UserId, &u.UserName, &u.UserEmail, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
