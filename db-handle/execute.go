package dbhandle

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Execute changes in the database.

func ExecuteDb(queryStr string) (rows sql.Result, err error) {

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return rows, err
	}
	defer db.Close()

	// makes alterations in the database
	rows, err = db.Exec(queryStr)

	return rows, err

}
