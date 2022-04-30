package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleLsGenre(limit uint) {

	// @TODO we can extract this to a function

	queryStr := "SELECT DISTINCT genre FROM movies ORDER BY genre ASC"
	if limit > 0 {
		queryStr += fmt.Sprintf(" LIMIT %d", limit)
	}
	rows, err := dbhandle.QueryDb(queryStr)

	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
	for rows.Next() {
		var genreLable string
		if err := rows.Scan(&genreLable); err != nil {
			log.Fatalf("An error occurred: %s \n", err.Error())
		}
		println(genreLable)
	}
}
