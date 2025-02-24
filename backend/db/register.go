package db

import (
	"fmt"
	strc "music/structs"
)

func WriteNewUser(new_user *strc.RegisterRequest) int {
	DB := DBConnection()
	defer DBDisconnection(DB)
	fmt.Println(new_user.BirthDate)
	_, err := DB.Exec("insert into User(username, email, password_hash,full_name, birth_date) values (?, ?, ?, ?, ?)", new_user.Username, new_user.Email, new_user.PasswordHash, new_user.FullName, new_user.BirthDate)
	if err != nil {
		fmt.Print(err)
		return 1
	}
	return 0
}
