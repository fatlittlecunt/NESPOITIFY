package db

import strc "music/structs"

func GetInfo() (*[]strc.ArtistForReport, error) {
	DB := DBConnection()
	defer DBDisconnection(DB)
	var artists []strc.ArtistForReport
	query := "Select artist_name, genre, popularity from artist"

	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a strc.ArtistForReport
		err := rows.Scan(&a.ArtistName, &a.Genre, &a.Popularity)
		if err != nil {
			return nil, err
		}
		artists = append(artists, a)
	}

	return &artists, nil
}
