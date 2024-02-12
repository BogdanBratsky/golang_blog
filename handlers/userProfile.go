package handlers

import (
	"blog/utils"
	"encoding/json"
	"log"
	"net/http"
)

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
