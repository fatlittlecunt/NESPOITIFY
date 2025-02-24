package db

import "errors"

func Unlike(song_id, user_id int) error {
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "delete from likes where userid = ? and songid = ?;"
	_, err := DB.Exec(query, user_id, song_id)
	if err != nil {
		return errors.New("delete error")
	}
	return nil
}
