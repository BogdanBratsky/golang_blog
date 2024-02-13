package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

// функция для генерации токена
func CreateToken(uId *uint64, uName *string) (string, error) {
	// Заголовок токена
	tokenHeader := jwt.New(jwt.SigningMethodHS256)

	// Полезная нагрузка токена
	claims := jwt.MapClaims{
		"sub":  *uId,                                       // Идентификатор пользователя
		"name": *uName,                                     // Имя пользователя
		"iat":  time.Now().Unix(),                          // Время выдачи токена (UNIX timestamp)
		"exp":  time.Now().Add(time.Hour * 24 * 14).Unix(), // Время истечения токена (UNIX timestamp)
	}

	tokenHeader.Claims = claims

	// Подпись токена
	tokenString, err := tokenHeader.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// функция для извлечения токена из запроса
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

func ParseToken(tokenString string) (uint64, error) {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Проверяем, является ли токен валидным
	if !token.Valid {
		return 0, errors.New("Invalid token")
	}

	// Извлекаем значение поля "sub" из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Failed to parse token claims")
	}

	subFloat, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("Failed to convert 'sub' to float64")
	}

	sub := uint64(subFloat)

	fmt.Println(sub)
	return sub, nil
}

func ValidateToken(tokenString, secretKey string) (*jwt.Token, error) {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем, является ли токен валидным
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return token, nil
}
