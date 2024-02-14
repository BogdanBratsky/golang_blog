package db

import "blog/models"

func GetUsersFromDB() ([]models.UserView, error) {
	// sql-запрос на получение данных из таблицы users
	rows, err := DB.Query("SELECT user_id, user_name, user_email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserView
	for rows.Next() {
		var u models.UserView
		err := rows.Scan(&u.UserId, &u.UserName, &u.UserEmail)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
