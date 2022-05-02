package cliinput

import (
	"database/sql"
	"fmt"
	"log"
)

// Print the resultant rows to the user.

func printIt(rows *sql.Rows) {
	count := 0
	for rows.Next() {
		var title string
		var id uint
		var year uint
		if err := rows.Scan(&id, &title, &year); err != nil {
			log.Fatalf("An error occurred: %s \n", err.Error())
		}
		count++
		fmt.Printf("%d \t %s [%d]\n", id, title, year)
	}
	if count == 0 {
		println("No record found for the given query")
	}

}
