package main

import (
	"blog/api/handlers"
	"blog/db"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/acme/autocert"
)

// функция main
func main() {

	db.InitDB()
	defer db.CloseDB()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// обработчики запросов
	r := mux.NewRouter()

	r.HandleFunc("/api/register", handlers.CreateUserHandler).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginUserHandler).Methods("POST")
	r.HandleFunc("/api/users", handlers.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", handlers.GetUserByIdHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}/articles", handlers.GetPostsByUserHandler).Methods("GET")
	r.HandleFunc("/api/articles", handlers.GetPostsHandler).Methods("GET")
	r.HandleFunc("/api/articles", handlers.CreatePostHandler).Methods("POST")
	r.HandleFunc("/api/articles/{id}", handlers.GetPostHandler).Methods("GET")
	r.HandleFunc("/api/articles/{id}", handlers.DeletePostHandler).Methods("DELETE")
	r.HandleFunc("/api/categories", handlers.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/api/categories/{id}", handlers.GetCategoryHandler).Methods("GET")
	r.HandleFunc("/api/categories/{id}/articles", handlers.GetPostsByCategoryHandler).Methods("GET")

	// Оборачиваем маршруты в обработчик CORS
	handlerWithCORS := c.Handler(r)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("alphaposts.online"), // Замените на свой домен
		Cache:      autocert.DirCache("certs"),                  // Кэширование сертификатов
	}

	server := &http.Server{
		Addr:      ":443",
		TLSConfig: &tls.Config{GetCertificate: certManager.GetCertificate},
		Handler:   handlerWithCORS,
	}

	// Запускаем HTTP-сервер для автоматического получения и обновления сертификатов
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))

	// Запускаем HTTPS-сервер
	err := server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}
}
