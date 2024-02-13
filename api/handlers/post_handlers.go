package handlers

import (
	"blog/db"
	"blog/models"
	"blog/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем токен из запроса
	tokenString := utils.ExtractToken(r)

	// Проверяем подлинность токена
	_, err := utils.ValidateToken(tokenString, "your-secret-key")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	var post models.Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}
	post.UserId, _ = utils.ParseToken(tokenString)
	post.CreatedAt = time.Now()

	if err = db.CreatePostDB(&post); err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	// Возвращаем успешный ответ с созданным пользователем
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем токен из запроса
	tokenString := utils.ExtractToken(r)

	// Проверяем подлинность токена
	_, err := utils.ValidateToken(tokenString, "your-secret-key")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	postId, _ := strconv.Atoi(vars["id"])
	userId, _ := utils.ParseToken(tokenString)

	result, err := db.DeletePostFromBD(&postId, &userId)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	if result != 1 {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "Article was deleted")
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
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

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
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
		Posts      []models.Post `json:"posts"`
		TotalCount int           `json:"totalCount"`
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

func GetPostsByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received:", r.Method, r.URL)

	vars := mux.Vars(r)
	categoryID, _ := strconv.Atoi(vars["id"])

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

	posts, totalCount, err := db.GetArticlesByCategoryDB(&categoryID, &page, &perPage)
	if err != nil {
		log.Println("Error fetching user's posts data:", err)
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	response := struct {
		Posts      []models.Post `json:"posts"`
		TotalCount int           `json:"totalCount"`
	}{
		Posts:      posts,
		TotalCount: totalCount,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}
}

func GetPostsByUserHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	response := struct {
		Posts      []models.Post `json:"posts"`
		TotalCount int           `json:"totalCount"`
	}{
		Posts:      posts,
		TotalCount: totalCount,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusUnauthorized)
		return
	}
}
