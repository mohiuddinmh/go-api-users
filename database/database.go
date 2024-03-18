package database

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=test dbname=test_db password=test sslmode=disable") // Replace connection details
	if err != nil {
		panic(err)
	}
}
