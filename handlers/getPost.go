package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL)

	vars := mux.Vars(r)
	postID, _ := strconv.Atoi(vars["id"])

	post, err := db.GetPostFromDB(&postID)
	if err != nil {
		http.Error(w, "Article doesn't exist", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(post)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
