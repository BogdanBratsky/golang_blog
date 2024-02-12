package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("your-secret-key")

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
