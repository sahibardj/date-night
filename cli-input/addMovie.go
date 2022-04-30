package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleAddMovie(title string, genre string, year uint) {

	// @TODO we can extzract this to a function

	queryStr := "INSERT INTO movies(title,genre,year) VALUES"
	queryStr += fmt.Sprintf("('%s','%s','%d');", title, genre, year)
	rows, err := dbhandle.ExecuteDb(queryStr)
	println(rows)
	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
}
