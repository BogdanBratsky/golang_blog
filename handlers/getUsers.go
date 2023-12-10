package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
)

// функция для обработчика, получающая все записи из таблицы users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// подключение к бд
	db.InitDB()
	defer db.CloseDB()

	// получение данных из бд
	users, err := db.GetUsersFromDB()
	if err != nil {
		log.Fatal(err)
	}

	// сериализация данных из бд
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
