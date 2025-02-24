package api

import (
	"encoding/json"
	"io"
	db "music/db"
	strc "music/structs"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// Убедимся, что метод POST
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
	var data strc.RegisterRequest
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	code := db.WriteNewUser(&data)
	if code == 0 {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(w, "Cannot create User. This username or email already exist", http.StatusForbidden)
		return
	}
}
