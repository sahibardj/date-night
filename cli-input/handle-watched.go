package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleLsWatched(limit uint) {

	queryStr := "SELECT id, title,year, genre FROM movies where watched=true ORDER BY title  ASC"
	if limit > 0 {
		queryStr += fmt.Sprintf(" LIMIT %d", limit)
	}
	rows, err := dbhandle.QueryDb(queryStr)

	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
	printIt(rows)
}
