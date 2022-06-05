package dbhandle

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func QueryDb(queryStr string) (rows *sql.Rows, err error) {

	// Open sqlite file and query the file using Query function call.

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return rows, err
	}
	defer db.Close()

	//use query funtion to fetch results.

	rows, err = db.Query(queryStr)

	return rows, err

}
