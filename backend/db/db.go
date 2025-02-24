package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (DB *sql.DB) {
	connStr := "root:root@/MusicService"
	DB, err := sql.Open("mysql", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connected")
	return DB
}

func DBDisconnection(DB *sql.DB) {
	DB.Close()
	fmt.Println("DB disconnected")
}

func GetUserID(email string) (int, error) {
	DB := DBConnection()
	defer DBDisconnection(DB)
	var id int
	query := "select userid from user where email = ?"
	rows, err := DB.Query(query, email)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("query error 1")
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return 0, errors.New("scanning error 1")

		}
	}
	return id, nil
}
