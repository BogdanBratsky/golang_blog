package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
)

func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL)

	categories, err := db.GetCategoriesFromDB()
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}

	// сериализация данных из бд
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}

	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Uncorrect request", http.StatusBadRequest)
		return
	}
}
