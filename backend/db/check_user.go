package db

import (
	"errors"
	"fmt"
	strc "music/structs"
)

func CheckUser(user *strc.LoginRequest) (string, error) {
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "Select role from User where email = ? and password_hash = ?"

	rows, err := DB.Query(query, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("query error 1")
	}
	defer rows.Close()

	var role string

	for rows.Next() {
		err := rows.Scan(&role)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("scanning error 1")

		}
	}
	fmt.Println(role)
	return role, nil

}
