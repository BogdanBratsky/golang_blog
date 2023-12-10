package db

func GetUsersFromDB() ([]User, error) {
	// sql-запрос на получение данных из таблицы users
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.UserId, &u.UserName, &u.UserEmail, &u.PasswordHash, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
