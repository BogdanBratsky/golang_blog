package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetArticlesByUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL)

	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

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

	posts, totalCount, err := db.GetArticlesByUserDB(&userID, &page, &perPage)
	if err != nil {
		log.Println("Error fetching user's posts data:", err)
		http.Error(w, "User doesn't exist", http.StatusBadRequest)
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
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
