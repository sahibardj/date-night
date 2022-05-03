package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleMovieWatched(idx uint) {


	queryStr := "UPDATE movies SET watched=true where id="
	queryStr += fmt.Sprintf("%d", idx)
	_, err := dbhandle.ExecuteDb(queryStr)

	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
}
