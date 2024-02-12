package utils

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

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
