package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDb(connstring string) error {
	var err error
	db, err = sql.Open("postgres", connstring)
	return err
}
