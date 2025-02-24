package api

import (
	"fmt"
	"music/db"
	"music/jwt"
	"net/http"
	"strconv"
	"strings"
)

func DeleteArtist(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	parts1 := strings.Split(authHeader, " ")
	if len(parts1) != 2 || parts1[0] != "Bearer" {
		http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
		return
	}

	claims, err := jwt.ParseAndVerifyToken(parts1[1])
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
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "ID артиста не указан", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID артиста должен быть числом", http.StatusBadRequest)
		return
	}

	err = db.DeleteArtist(id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
}
