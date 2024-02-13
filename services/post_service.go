package services

import (
	"blog/db"
	"blog/models"
	"blog/utils"
	"encoding/json"
	"time"
)

func CreatePost(tokenString string, requestBody *json.Decoder) (*models.Post, error) {
	// Проверяем подлинность токена
	userId, err := utils.ValidateToken(tokenString, "your-secret-key")
	if err != nil {
		return nil, err
	}

	// Декодируем тело запроса
	var post models.Post
	err = requestBody.Decode(&post)
	if err != nil {
		return nil, err
	}

	// Устанавливаем значения UserId и CreatedAt
	post.UserId = userId
	post.CreatedAt = time.Now()

	// Вызываем метод создания поста в базе данных
	err = db.CreatePostDB(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
