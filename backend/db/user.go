package db

import (
	"errors"
	"fmt"
	strc "music/structs"
)

func GetUserByID(id int) (*strc.User, error) {
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "select userid, username, email, full_name, birth_date, profile_picture_url, role from user where userid = ?"

	rows, err := DB.Query(query, id)
	if err != nil {
		return nil, errors.New("query error")
	}
	defer rows.Close()
	var user strc.User
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.FullName, &user.BirthDate, &user.ProfilePictureURL, &user.Role)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("scanning errorrrr")
		}

	}
	return &user, nil
}
