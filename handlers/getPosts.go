package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// функция для обработчика, получающая все записи из таблицы posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL)

	pageStr := r.FormValue("page")
	perPageStr := r.FormValue("perPage")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		perPage = 30
	}

	posts, totalCount, err := db.GetPostsFromDB(page, perPage)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}

	response := struct {
		Posts      []db.Post `json:"posts"`
		TotalCount int       `json:"totalCount"`
	}{
		Posts:      posts,
		TotalCount: totalCount,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}
}
