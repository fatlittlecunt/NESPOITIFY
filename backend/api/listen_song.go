package api

import (
	"encoding/json"
	"io"
	"music/db"
	strc "music/structs"
	"net/http"
	"os"
)

func ListenSong(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	// Парсим JSON
	var data strc.ListenSongRequest
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	songID := data.SongID
	file_name, err := db.GetFileName(songID)
	if err != nil {
		http.Error(w, "Cannot find this song", http.StatusInternalServerError)
		return
	}
	folder_path := "D:\\res\\music\\"
	audioFile, err := os.Open(folder_path + file_name)
	if err != nil {
		http.Error(w, "Audio file not found.", http.StatusNotFound)
		return
	}
	defer audioFile.Close()

	// Устанавливаем заголовки для аудио стрима
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Transfer-Encoding", "chunked")

	// Буфер для чтения файла
	buf := make([]byte, 1024)

	// Читаем и передаем данные по частям
	for {
		n, err := audioFile.Read(buf)
		if err != nil {
			break
		}
		if n > 0 {
			_, err := w.Write(buf[:n])
			if err != nil {
				break
			}
		}

		// Небольшая задержка, чтобы имитировать стриминг
		//time.Sleep(100 * time.Millisecond)
	}
}
