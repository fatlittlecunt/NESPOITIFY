package db

import (
	"errors"
	"fmt"
	strc "music/structs"

	_ "github.com/go-sql-driver/mysql"
)

func GetArtists() (*strc.ArtistsResponse, error) {
	var artists []strc.ArtistResponse
	DB := DBConnection()
	defer DBDisconnection(DB)

	query := "Select * from Artist"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, errors.New("query error")
	}
	defer rows.Close()

	for rows.Next() {
		var artist strc.ArtistResponse
		err := rows.Scan(&artist.ArtistID, &artist.ArtistName, &artist.Country, &artist.Genre, &artist.Description, &artist.Popularity, &artist.ArtistPicture)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("scanning errorrrr")
		}
		artists = append(artists, artist)
	}

	result := strc.ArtistsResponse{
		ArtistsList: artists,
	}

	return &result, nil
}
