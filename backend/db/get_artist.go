package db

import (
	"fmt"
	strc "music/structs"

	"errors"

	_ "github.com/go-sql-driver/mysql"
)

func GetArtist(id int) (*strc.GetArtistResponse, error) {

	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "Select * from Artist where ArtistID = ?"

	rows, err := DB.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("query error 1")
	}
	defer rows.Close()

	var artist strc.ArtistResponse

	for rows.Next() {
		err := rows.Scan(&artist.ArtistID, &artist.ArtistName, &artist.Country, &artist.Genre, &artist.Description, &artist.Popularity, &artist.ArtistPicture)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("scanning error 1")

		}
	}

	query = "select SongID, title, duration, song_popularirty, song_picture_url, file_url from SongsArtists where ArtistID = ?"

	rows, err = DB.Query(query, artist.ArtistID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("query error 2")
	}
	defer rows.Close()

	var songs []strc.SongsForArtists
	for rows.Next() {
		var song strc.SongsForArtists
		err := rows.Scan(&song.SongID, &song.SongTitle, &song.Duration, &song.SongPopularity, &song.Picture, &song.SongURL)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("scanning error 2")
		}
		songs = append(songs, song)
	}

	result := strc.GetArtistResponse{
		Artist: artist,
		Songs:  songs,
	}

	return &result, nil
}
