package handlers

import (
	"blog/db"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL)

	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	// db.InitDB()
	// defer db.CloseDB()

	user, err := db.GetUserFromDB(&userID)
	if err != nil {
		log.Println("Error fetching user data:", err)
		http.Error(w, "User doesn't exist", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}
}
