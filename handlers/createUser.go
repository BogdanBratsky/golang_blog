package handlers

import (
	"blog/db"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Парсинг JSON из тела запроса
	var user db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Открываем соединение с базой данных
	db.InitDB()
	// Закрываем соединение с базой данных
	defer db.CloseDB()

	// Проверка на оригинальность почты
	res, _ := db.EmailExists(&user.UserEmail)

	if !res {
		// SQL-запрос для вставки пользователя
		if err := db.CreateUserDB(&user); err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		// Возвращаем успешный ответ с созданным пользователем
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	} else {
		http.Error(w, "This email is busy", http.StatusInternalServerError)
		return
	}
}
