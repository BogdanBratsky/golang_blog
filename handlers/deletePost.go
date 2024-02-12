package handlers

import (
	"blog/db"
	"blog/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Извлекаем токен из запроса
	tokenString := utils.ExtractToken(r)

	// Проверяем подлинность токена
	_, err := utils.ValidateToken(tokenString, "your-secret-key")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	postId, _ := strconv.Atoi(vars["id"])
	userId, _ := utils.ParseToken(tokenString)

	result, err := db.DeletePostFromBD(&postId, &userId)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}

	if result != 1 {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, "Article was deleted")
}
