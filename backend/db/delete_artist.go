package db

import (
	"errors"
	"fmt"
)

func DeleteArtist(id int) error {
	DB := DBConnection()
	defer DBDisconnection(DB)

	fmt.Println(id)

	query1 := "delete from songartist where artistid = ?"
	query2 := "delete from artist where artistid = ?"

	_, err := DB.Exec(query1, id)
	if err != nil {
		return errors.New("query error 1")
	}

	_, err = DB.Exec(query2, id)
	if err != nil {
		return errors.New("query error 2")
	}
	return nil
}
