package main

import (
	"blog/db"
	"blog/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// функция main
func main() {

	// обработчики запросов
	r := mux.NewRouter()

	r.HandleFunc("/profile", handlers.UserProfileHandler).Methods("GET")
	r.HandleFunc("/register", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/login", handlers.SignInUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/users/{id}/articles", handlers.GetArticlesByUser).Methods("GET")
	r.HandleFunc("/articles", handlers.GetPosts).Methods("GET")
	r.HandleFunc("/articles", handlers.CreatePostHandler).Methods("POST")
	r.HandleFunc("/articles/{id}", handlers.GetPost).Methods("GET")
	r.HandleFunc("/articles/{id}", handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/categories", handlers.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/categories/{id}", handlers.GetCategoryHandler).Methods("GET")
	r.HandleFunc("/categories/{id}/articles", handlers.GetArticlesByCategory).Methods("GET")

	db.InitDB()
	defer db.CloseDB()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Оборачиваем маршруты в обработчик CORS
	handlerWithCORS := c.Handler(r)

	// Запускаем сервер
	err := http.ListenAndServe(":3000", handlerWithCORS)
	if err != nil {
		log.Fatal(err)
	}
}
