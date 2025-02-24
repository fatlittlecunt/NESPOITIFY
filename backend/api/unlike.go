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

func Unlike(w http.ResponseWriter, r *http.Request) {
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
	email := claims["sub"].(string)
	userid, err := db.GetUserID(email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Denied", http.StatusUnauthorized)
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
		fmt.Println(songid)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	fmt.Println(songid.SongID)
	song_id, err := strconv.Atoi(songid.SongID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("BLABLA")
		http.Error(w, "convert error", http.StatusInternalServerError)
	}

	err = db.Unlike(song_id, userid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "unlike error", http.StatusInternalServerError)
	}
}
