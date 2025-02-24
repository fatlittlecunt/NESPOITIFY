package db

import (
	"errors"
	"fmt"
	t "music/tool"
	"os"
	"strconv"
	"strings"
)

func EditArtist(id int, name, country, genre, desc, pic string) error {
	DB := DBConnection()
	defer DBDisconnection(DB)

	_, err := DB.Exec("update artist set artist_name = ? where artistid = ?", name, id)
	if err != nil {
		return errors.New("query error 1")
	}

	_, err = DB.Exec("update artist set country = ? where artistid = ?", country, id)
	if err != nil {
		return errors.New("query error 2")
	}

	_, err = DB.Exec("update artist set genre = ? where artistid = ?", genre, id)
	if err != nil {
		return errors.New("query error 3")
	}

	_, err = DB.Exec("update artist set artist_description = ? where artistid = ?", desc, id)
	if err != nil {
		return errors.New("query error 4")
	}

	if len(pic) < 70 {
		return nil
	}

	pic_path := "D:/res/artist_profile_photo/" + strconv.Itoa(id) + ".png"

	if _, err := os.Stat(pic_path); err == nil {
		// Файл существует, удаляем его
		err := os.Remove(pic_path)
		if err != nil {
			return fmt.Errorf("не удалось удалить существующий файл: %w", err)
		}
	}
	var img *[]byte
	if strings.HasPrefix(pic, "data:image") {
		img, err = t.ParseBase64Image(pic)
		if err != nil {
			return errors.New("image procces error")
		}
	}

	if !strings.HasPrefix(pic, "http") {
		outputPath := "D:/res/artist_profile_photo"
		file := outputPath + "/" + strconv.Itoa(id) + ".png"
		fmt.Println(file)
		err = os.WriteFile(file, *img, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return errors.New("cannot save image")
		}
	}

	artist_picture_url := "http://localhost:8080/images/band/" + strconv.Itoa(id) + ".png"

	_, err = DB.Exec("update artist set artist_picture_url = ? where artistid = ?", artist_picture_url, id)
	if err != nil {
		return errors.New("query error 5")
	}

	return nil
}
