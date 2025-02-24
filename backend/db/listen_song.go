package db

import (
	"errors"
	"fmt"
	"strings"
)

func GetFileName(song_id int) (string, error) {
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "Select file_url from Song where SongID = ?"

	rows, err := DB.Query(query, song_id)
	if err != nil {
		fmt.Println(err)
		return "error 1", errors.New("query error 1")
	}
	defer rows.Close()

	var path string

	for rows.Next() {
		err := rows.Scan(&path)
		if err != nil {
			fmt.Println(err)
			return "error 2", errors.New("scanning error 2")

		}
	}
	path = strings.TrimPrefix(path, "\n")
	path = strings.TrimPrefix(path, "\r")

	return path, nil
}
