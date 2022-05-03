package cliinput

import (
	"fmt"
	"log"
	dbhandle "movie_database/db-handle"
)

func handleRemoveById(idx uint) {


	queryStr := "DELETE FROM movies where id="
	queryStr += fmt.Sprintf("%d", idx)
	_, err := dbhandle.ExecuteDb(queryStr)
	
	if err != nil {
		log.Fatalf("An error occurred: %s \n", err.Error())
	}
}
