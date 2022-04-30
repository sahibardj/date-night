package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleSurpriseMe(genreLable string) {

	// @TODO we can extract this to a function

	queryStr := "SELECT id, title,year FROM movies where watched=false AND "
	queryStr += fmt.Sprintf("genre='%s'", genreLable)
	queryStr += " ORDER BY RANDOM()"

	rows, err := dbhandle.QueryDb(queryStr)

	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
	for rows.Next() {
		var title string
		var id uint
		var year uint
		if err := rows.Scan(&id, &title, &year); err != nil {
			log.Fatalf("An error occurred: %s \n", err.Error())
		}
		fmt.Printf("%d \t %s [%d]\n", id, title, year)
	}
}
