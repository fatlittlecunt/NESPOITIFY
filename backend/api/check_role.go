package api

import (
	"fmt"
	"music/jwt"
	"net/http"
	"strings"
)

func CheckRole(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
		return
	}

	claims, err := jwt.ParseAndVerifyToken(parts[1])
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Denied", http.StatusUnauthorized)
		return
	}

	// Извлекаем данные из claims
	role := claims["role"].(string)
	if role != "manager" {
		http.Error(w, "Not manager", http.StatusForbidden)
	}
}
