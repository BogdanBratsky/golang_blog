package utils

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

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
