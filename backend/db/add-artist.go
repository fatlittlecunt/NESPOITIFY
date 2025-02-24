package db

import (
	"errors"
	"fmt"
	t "music/tool"
	"os"
	"strconv"
	"strings"
)

func AddArtist(artist_name, country, genre, desc, pic string) error {
	fmt.Println("ENTER ADDARTIST")
	DB := DBConnection()
	defer DBDisconnection(DB)
	var img *[]byte
	var err error
	var name string
	var artist_id int
	query := "Select * from Artist where artist_name = ?"
	rows, err := DB.Query(query, artist_name)
	if err != nil {
		return errors.New("query error 1")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			return errors.New("query error 2")
		}
	}

	if name != "" {
		return errors.New("exists")
	}

	query = "INSERT INTO Artist (artist_name, country, genre, artist_description) values (?, ?, ?, ?)"

	_, err = DB.Exec(query, artist_name, country, genre, desc)
	if err != nil {
		return errors.New("query error 3")
	}

	if strings.HasPrefix(pic, "data:image") {
		img, err = t.ParseBase64Image(pic)
		if err != nil {
			return errors.New("image procces error")
		}
	}

	query = "Select ArtistID from Artist where artist_name = ?"
	rows, err = DB.Query(query, artist_name)
	if err != nil {
		return errors.New("query error 4")
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&artist_id)
		if err != nil {
			return errors.New("query error 5")
		}
	}

	if len(pic) < 30 {
		return nil
	}

	if !strings.HasPrefix(pic, "http") {
		fmt.Println("ГОЙДАААААААААААААААААА")
		outputPath := "D:/res/artist_profile_photo"
		file := outputPath + "/" + strconv.Itoa(artist_id) + ".png"
		fmt.Println(file)
		err = os.WriteFile(file, *img, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return errors.New("cannot save image")
		}

	}

	query = "update Artist set artist_picture_url = ? where artistid = ?"
	path := "http://localhost:8080/images/band/" + strconv.Itoa(artist_id) + ".png"
	fmt.Println(path)

	_, err = DB.Exec(query, path, artist_id)
	if err != nil {
		fmt.Println(err)
		return errors.New("query error 4")
	}
	return nil
}
