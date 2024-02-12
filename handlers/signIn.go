package handlers

import (
	"blog/db"
	"blog/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func SignInUser(w http.ResponseWriter, r *http.Request) {
	var user db.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// db.InitDB()
	// defer db.CloseDB()

	res, _ := db.UserExists(&user.UserEmail, &user.PasswordHash)

	if res > 0 {
		fmt.Println("User exists")
		// Пользователь существует, создаем JWT токен
		token, err := utils.CreateToken(&res, &user.UserName)
		if err != nil {
			http.Error(w, "Error creating token", http.StatusInternalServerError)
			fmt.Println("Error creating token:", err)
			return
		}

		// Отправляем токен в ответе
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	} else {
		fmt.Println("User doesn't exist")
		// Пользователь не найден
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
