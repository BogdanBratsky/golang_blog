package handlers

import (
	"blog/db"
	"encoding/json"
	"net/http"
)

// функция для обработчика, получающая все записи из таблицы users
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// получение данных из бд
	users, err := db.GetUsersFromDB()
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}

	// сериализация данных из бд
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}
}
