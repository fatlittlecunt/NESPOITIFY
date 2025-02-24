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

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
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
	var edit strc.UpdateProfileRequest
	if err := json.Unmarshal(body, &edit); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	err = db.EditProfile(edit.Username, edit.FullName, edit.ProfilePicture, userid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Edit error", http.StatusInternalServerError)
	}
}
