package handlers

import (
	"blog/db"
	"blog/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем токен из запроса
	tokenString := utils.ExtractToken(r)

	// Проверяем подлинность токена
	_, err := utils.ValidateToken(tokenString, "your-secret-key")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var post db.Post

	db.InitDB()
	defer db.CloseDB()

	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.CreatePostDB(&post); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ с созданным пользователем
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
