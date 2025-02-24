package api

import (
	"encoding/json"
	"io"
	"music/db"
	jwtgen "music/jwt"
	strc "music/structs"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")              // Разрешаем запросы с любых источников
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") // Разрешаем методы
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Читаем тело запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	// Парсим JSON
	var data strc.LoginRequest
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	role, err := db.CheckUser(&data)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	if role == "" {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	JWT, err := jwtgen.JWTGen(&data, role)
	if err != nil {
		http.Error(w, "Cannot generate token", http.StatusInternalServerError)
	}
	response := strc.LoginResponse{
		AccessToken: JWT,
	}

	w.Header().Set("Content-Type", "application/json")

	// Кодируем ответ в JSON и отправляем его
	json.NewEncoder(w).Encode(response)
}
