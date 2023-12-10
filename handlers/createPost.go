package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {

	var post db.Post

	db.InitDB()
	defer db.CloseDB()

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.CreatePostDB(&post); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ с созданным пользователем
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)

	// // SQL-запрос для вставки пользователя
	// _, err = db.Exec(`INSERT INTO posts (post_title, post_content) VALUES ($1, $2) RETURNING post_id`,
	// 	post.PostTitle, post.PostContent)
	// if err != nil {
	// 	http.Error(w, "Failed to create user", http.StatusInternalServerError)
	// 	return
	// }

}
