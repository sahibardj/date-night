package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)
//Add movie to the database
func handleAddMovie(title string, genre string, year uint) {


	queryStr := "INSERT INTO movies(title,genre,year) VALUES"
	queryStr += fmt.Sprintf("('%s','%s','%d');", title, genre, year)
	_, err := dbhandle.ExecuteDb(queryStr)
	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
}
