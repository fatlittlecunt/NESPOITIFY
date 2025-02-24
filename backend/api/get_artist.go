package api

import (
	"encoding/json"
	"fmt"
	db "music/db"
	"net/http"
	"strconv"
	"strings"
)

func GetArtist(w http.ResponseWriter, r *http.Request) {

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

	artists, err := db.GetArtist(id)
	if err != nil {
		http.Error(w, "Cannot get artists", http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	// Кодируем ответ в JSON и отправляем его
	json.NewEncoder(w).Encode(artists)
}
