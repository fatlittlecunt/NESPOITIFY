package api

import (
	"encoding/json"
	"fmt"
	db "music/db"
	"net/http"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {

	artists, err := db.GetArtists()
	if err != nil {
		http.Error(w, "Cannot get artists", http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	// Кодируем ответ в JSON и отправляем его
	json.NewEncoder(w).Encode(artists)
}
