package db

import (
	"errors"
	"fmt"
)

func DeleteSong(id int) error {
	DB := DBConnection()
	defer DBDisconnection(DB)

	fmt.Println(id)

	query1 := "delete from song where songid = ?"

	_, err := DB.Exec(query1, id)
	if err != nil {
		return errors.New("query error 1")
	}

	return nil
}
