package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleLsWatched(limit uint) {

	// @TODO we can extract this to a function

	queryStr := "SELECT title  FROM movies where watched=true ORDER BY title  ASC"
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
