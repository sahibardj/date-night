package cliinput

import (
	"fmt"
	"database/sql"
	"log"
)

func printIt(rows *sql.Rows){

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
