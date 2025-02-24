package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"music/db"
	"music/jwt"
	strc "music/structs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetReport(w http.ResponseWriter, r *http.Request) {
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

	err = db.CreateReport()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "cannot create report", http.StatusInternalServerError)
	}

	outputPath := filepath.Join("D:", "res", "report.xlsx")

	// Открываем файл
	file, err := os.Open(outputPath)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Ошибка открытия файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка открытия файла: "+err.Error(), http.StatusInternalServerError)
	}

	// Кодируем содержимое файла в Base64
	base64Str := base64.StdEncoding.EncodeToString(fileData)
	request := strc.Base64Excel{
		Base64: base64Str,
	}

	w.Header().Set("Content-Type", "application/json")

	// Кодируем ответ в JSON и отправляем его
	json.NewEncoder(w).Encode(request)

}
