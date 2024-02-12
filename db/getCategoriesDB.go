package db

func GetCategoriesFromDB() ([]Category, error) {
	// sql-запрос на получение данных из таблицы users
	rows, err := DB.Query("SELECT category_id, category_name FROM categories ORDER BY category_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		err := rows.Scan(&c.CategoryId, &c.CategoryName)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}
