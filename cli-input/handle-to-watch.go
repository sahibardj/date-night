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
	for rows.Next() {
		var genreLable string
		var title string
		var id uint
		var year uint
		if err := rows.Scan(&id, &title, &year, &genreLable); err != nil {
			log.Fatalf("An error occurred: %s \n", err.Error())
		}
		fmt.Printf("%d \t %s [%d]\t%s\n", id, title, year, genreLable)
	}
}
