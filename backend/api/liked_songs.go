package api

import (
	"encoding/json"
	"fmt"
	"music/db"
	"music/jwt"
	"net/http"
	"strings"
)

func LikedSongs(w http.ResponseWriter, r *http.Request) {
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
	email := claims["sub"].(string)
	user_id, err := db.GetUserID(email)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
	}

	songs, err := db.GetLiked(user_id)
	if err != nil {
		http.Error(w, "not found", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")

	// Кодируем ответ в JSON и отправляем его
	json.NewEncoder(w).Encode(songs)
}
