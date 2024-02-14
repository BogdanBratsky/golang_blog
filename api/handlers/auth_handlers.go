package handlers

import (
	"blog/db"
	"blog/models"
	"blog/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Парсинг JSON из тела запроса
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	// Проверка на оригинальность почты
	res, _ := db.EmailExists(&user.UserEmail)

	if !res {
		// SQL-запрос для вставки пользователя
		if err := db.CreateUserDB(&user); err != nil {
			http.Error(w, "Internal Server Error", http.StatusUnauthorized)
			return
		}

		// Возвращаем успешный ответ с созданным пользователем
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	} else {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	type result struct {
		UserInfo models.UserView `json:"userInfo"`
		Token    string          `json:"token"`
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id, name, _ := db.UserExists(&user.UserEmail, &user.PasswordHash)

	if id != 0 {
		fmt.Println("User exists")
		// Пользователь существует, создаем JWT токен
		token, err := utils.CreateToken(&id, &name)
		if err != nil {
			http.Error(w, "Error creating token", http.StatusInternalServerError)
			return
		}

		userView := models.UserView{
			UserId:    id,
			UserName:  name,
			UserEmail: user.UserEmail,
		}

		// Отправляем токен в ответе
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result{userView, token})
	} else {
		http.Error(w, "User doesn't exist", http.StatusUnauthorized)
		return
	}
}
