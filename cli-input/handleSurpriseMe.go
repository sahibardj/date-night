package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleSurpriseMe(genreLable string) {

	queryStr := "SELECT id, title,year FROM movies where watched=false AND "
	queryStr += fmt.Sprintf("genre='%s'", genreLable)
	queryStr += " ORDER BY RANDOM() "
	queryStr += fmt.Sprintf("LIMIT %d", 1)

	rows, err := dbhandle.QueryDb(queryStr)

	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
	//parse through the resultant rows and display output.
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
