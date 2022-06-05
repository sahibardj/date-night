package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

//List all the genre
func handleLsGenre(limit uint) {

	// Save the query to be done on the database in a string.
	queryStr := "SELECT DISTINCT genre FROM movies ORDER BY genre ASC"
	if limit > 0 {
		queryStr += fmt.Sprintf(" LIMIT %d", limit)
	}
	//pass the query to the query function (fetches/reads data from the database.
	rows, err := dbhandle.QueryDb(queryStr)

	//check for error
	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
	// Parse through the resultant and display output to the command line.
	for rows.Next() {
		var genreLable string
		if err := rows.Scan(&genreLable); err != nil {
			log.Fatalf("An error occurred: %s \n", err.Error())
		}
		println(genreLable)
	}
}
