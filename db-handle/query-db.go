package dbhandle

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func QueryDb(queryStr string) (rows *sql.Rows, err error) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return rows, err
	}
	defer db.Close()

	rows, err = db.Query(queryStr)

	return rows, err
}
