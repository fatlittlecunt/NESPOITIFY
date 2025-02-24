package jwt

import (
	"errors"
	"fmt"
	"log"
	strc "music/structs"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var StopSlovo = []byte("very-secret-key")

func JWTGen(user *strc.LoginRequest, role string) (string, error) {
	payload := jwt.MapClaims{
		"sub":  user.Email,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	t, err := token.SignedString(StopSlovo)
	if err != nil {
		log.Printf("JWT gen error")
		return "", errors.New("JWT gen error")
	}
	return t, nil
}

func ParseAndVerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Парсинг токена и проверка подписи
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Убедимся, что используется метод подписи HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return StopSlovo, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, errors.New("invalid token")
	}

	// Проверяем, что токен валиден
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Printf("Invalid token claims: %v", err)
		return nil, errors.New("invalid token claims")
	}
}
