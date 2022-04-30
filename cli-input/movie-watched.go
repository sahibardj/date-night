package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleMovieWatched(idx uint) {

	// @TODO we can extzract this to a function

	queryStr := "UPDATE movies SET watched=true where id="
	queryStr += fmt.Sprintf("%d", idx)
	rows, err := dbhandle.ExecuteDb(queryStr)
	println(rows)
	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
}
