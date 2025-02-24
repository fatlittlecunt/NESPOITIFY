package db

import (
	"errors"
	"fmt"
	strc "music/structs"
	t "music/tool"
	"os"
	"strconv"
	"strings"
)

func AddSong(artist_id int, song *strc.AddSongRequest) error {
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "insert into song (title, song_genre, duration, release_date, song_language, file_url) values (?, ?, ?, ?, ?, ?)"

	_, err := DB.Exec(query, song.SongName, song.SongGenre, song.SongDuration, song.SongDate, song.SongLanguage, "")
	if err != nil {
		fmt.Println(err)
		return errors.New("query 1 error")
	}

	query = "select songid from song where title = ? and song_genre = ? and duration = ? and release_date = ?"
	rows, err := DB.Query(query, song.SongName, song.SongGenre, song.SongDuration, song.SongDate)
	if err != nil {
		fmt.Println(err)
		return errors.New("query 2 error")
	}
	defer rows.Close()
	var song_id int
	for rows.Next() {
		err := rows.Scan(&song_id)
		if err != nil {
			fmt.Println(err)
			return errors.New("query error 3")
		}
	}
	var aud *[]byte
	if strings.HasPrefix(song.SongFile, "data:audio") {
		aud, err = t.ParseBase64Image(song.SongFile)
		if err != nil {
			return errors.New("file procces error")
		}
	}

	if !strings.HasPrefix(song.SongFile, "http") {
		outputPath := "D:/res/music"
		file := outputPath + "/" + strconv.Itoa(song_id) + ".mp3"
		fmt.Println(file)
		err = os.WriteFile(file, *aud, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return errors.New("cannot save songfile")
		}

	}

	query = "insert into songartist (SongID, ArtistID) values (?, ?)"
	_, err = DB.Exec(query, song_id, song.ArtistID)
	if err != nil {
		fmt.Println(err)
		return errors.New("query 6 error")
	}

	path := "http://localhost:8080/music/" + strconv.Itoa(song_id) + ".mp3"
	query = "update song set file_url = ? where songid = ?"
	_, err = DB.Exec(query, path, song_id)
	if err != nil {
		return errors.New("query 7 error")
	}

	return nil
}
