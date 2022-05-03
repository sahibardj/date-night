package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleLsToWatch(limit uint) {

	// @TODO we can extract this to a function

	queryStr := "SELECT id, title,year, genre FROM movies where watched=false ORDER BY title  ASC"
	if limit > 0 {
		queryStr += fmt.Sprintf(" LIMIT %d", limit)
	}
	rows, err := dbhandle.QueryDb(queryStr)

	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
	printIt(rows)
}
