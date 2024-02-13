package db

import (
	"blog/models"
	"database/sql"
	"log"
)

func GetCategoryFromDB(id *int) (models.Category, error) {
	var c models.Category
	query := `SELECT * FROM categories WHERE category_id = $1`
	err := DB.QueryRow(query, *id).Scan(&c.CategoryId, &c.CategoryName)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, err
		}
		log.Println("Error executing query:", err)
		return c, err
	}
	return c, err
}
