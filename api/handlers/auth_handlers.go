package handlers

import (
	"blog/db"
	"blog/models"
	"blog/utils"
	"encoding/json"
	"fmt"
	"log"
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
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

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

func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем токен из запроса
	tokenString := utils.ExtractToken(r)

	// Проверяем подлинность токена
	_, err := utils.ValidateToken(tokenString, "your-secret-key")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userId, _ := utils.ParseToken(tokenString)

	jsonData, err := json.Marshal(userId)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
