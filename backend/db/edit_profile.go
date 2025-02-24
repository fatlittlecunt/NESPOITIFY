package db

import (
	"errors"
	"fmt"
	t "music/tool"
	"os"
	"strconv"
	"strings"
)

func EditProfile(new_username, new_fullname, new_pic string, user_id int) error {
	DB := DBConnection()
	defer DBDisconnection(DB)
	var img *[]byte
	var err error
	if strings.HasPrefix(new_pic, "data:image") {
		img, err = t.ParseBase64Image(new_pic)
		if err != nil {
			return errors.New("image procces error")
		}
	}

	if !strings.HasPrefix(new_pic, "http") {
		outputPath := "D:/res/user_profile_photo"
		file := outputPath + "/" + strconv.Itoa(user_id) + ".png"
		fmt.Println(file)
		err = os.WriteFile(file, *img, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return errors.New("cannot save image")
		}
		query := "update user set profile_picture_url = replace(profile_picture_url, \"default\", ?) where userid = ?"
		_, err = DB.Exec(query, user_id, user_id)
		if err != nil {
			fmt.Print(err)
			return nil
		}
	}

	query := "update user set username = ? where userid = ?"
	_, err = DB.Exec(query, new_username, user_id)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	query = "update user set full_name = ? where userid = ?"
	_, err = DB.Exec(query, new_fullname, user_id)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	return nil
}
