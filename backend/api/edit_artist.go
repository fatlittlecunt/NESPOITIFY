package api

import (
	"encoding/json"
	"fmt"
	"io"
	"music/db"
	"music/jwt"
	strc "music/structs"
	"net/http"
	"strings"
)

func EditArtist(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		fmt.Println("1")
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		fmt.Println("2")
		http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
		return
	}
	///працюэ
	claims, err := jwt.ParseAndVerifyToken(parts[1])
	if err != nil {
		fmt.Println(err, "1")
		http.Error(w, "Denied", http.StatusUnauthorized)
		return
	}

	// Извлекаем данные из claims
	role := claims["role"].(string)

	if role != "manager" {
		http.Error(w, "Denied", http.StatusUnauthorized)
		return
	}
	//было
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "2")
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	// Парсим JSON
	var artist strc.ArtistResponse
	if err := json.Unmarshal(body, &artist); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	err = db.EditArtist(artist.ArtistID, artist.ArtistName, artist.Country, artist.Genre, artist.Description, artist.ArtistPicture)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "DB error", http.StatusInternalServerError)
	}
}
