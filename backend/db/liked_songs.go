package db

import (
	"errors"
	"fmt"
	strc "music/structs"
)

func GetLiked(user_id int) (*[]strc.SongsForArtists, error) {
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "select s.SongID, s.title, s.song_popularirty, s.duration, s.song_picture_url, s.file_url from song s join userslikes u on u.songid = s.songid where u.userid = ? and u.isliked = 1;"

	rows, err := DB.Query(query, user_id)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("query error 1")
	}
	defer rows.Close()

	var songs []strc.SongsForArtists

	for rows.Next() {
		var song strc.SongsForArtists
		err := rows.Scan(&song.SongID, &song.SongTitle, &song.SongPopularity, &song.Duration, &song.Picture, &song.SongURL)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("scanning error 1")
		}
		songs = append(songs, song)
	}
	return &songs, nil
}
