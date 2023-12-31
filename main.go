package main

import (
	"blog/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// функция main
func main() {
	// обработчики запросов
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/posts", handlers.CreatePostHandler).Methods("POST")
	// r.HandleFunc("/register", handlers.CreateUser).Methods("POST")
	// r.HandleFunc("/login", handlers.signInUser).Methods("POST")

	// запуск сервера
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
