package db

func LikeSong(user_id, song_id int) error {
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "insert into likes (userid, songid) values (?, ?)"
	_, err := DB.Exec(query, user_id, song_id)
	if err != nil {
		return err
	}
	return nil
}
