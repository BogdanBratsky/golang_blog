package utils

import (
	"net/http"
	"strings"
)

func ExtractToken(r *http.Request) string {
	// Получаем значение заголовка Authorization
	authHeader := r.Header.Get("Authorization")

	// Проверяем, что заголовок существует и начинается с "Bearer "
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}

	// Извлекаем токен из строки "Bearer <token>"
	return strings.TrimPrefix(authHeader, "Bearer ")
}
