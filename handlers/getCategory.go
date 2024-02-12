package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCategoryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL)

	vars := mux.Vars(r)
	categoryID, _ := strconv.Atoi(vars["id"])

	// db.InitDB()
	// defer db.CloseDB()

	category, err := db.GetCategoryFromDB(&categoryID)
	if err != nil {
		log.Println("Error fetching category data:", err)
		http.Error(w, "Category doesn't exist", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(category)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
