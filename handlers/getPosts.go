package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
)

// функция для обработчика, получающая все записи из таблицы posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	// подключение к бд
	db.InitDB()
	defer db.CloseDB()

	posts, err := db.GetPostsFromDB()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(posts)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
