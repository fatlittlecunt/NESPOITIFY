package api

import (
	"encoding/json"
	"fmt"
	"io"
	"music/db"
	"music/jwt"
	strc "music/structs"
	"net/http"
	"strconv"
	"strings"
)

func LikeSong(w http.ResponseWriter, r *http.Request) {
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
	role := claims["role"].(string)

	if role == "no_subscribe" {
		http.Error(w, "Denied", http.StatusForbidden)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	// Парсим JSON
	var songid strc.TrackPlay
	if err := json.Unmarshal(body, &songid); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	song_id, err := strconv.Atoi(songid.SongID)
	if err != nil {
		http.Error(w, "convert error", http.StatusInternalServerError)
	}

	user_id, err := db.GetUserID(email)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
	}

	err = db.LikeSong(user_id, song_id)
	if err != nil {
		http.Error(w, "like error", http.StatusInternalServerError)
	}

}
