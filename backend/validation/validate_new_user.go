package validation

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ValidateUsername(DB *sql.DB, username string) {
	rows, err := DB.Query("select count(*) from MusicService.User where username = ?", username)
	if err != nil {
		log.Println(err)
	}

	fmt.Print(rows)
}
