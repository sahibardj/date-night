package cliinput

import "fmt"

func printHelp() {
	println("WELCOME TO DATE NIGHT MOVIE SELECTOR")

	fmt.Printf("\n\n")

	println("ls-genre\n")
	println("\t\t To list out all the genres available and can have a limit to  number of genre listed. append '-limit=(a number here)'")
	fmt.Printf("\n\n")

	println("ls-towatch\n")
	println("\t\tCommand to have all movies you have to watch. To have a limited number of listing append '-limit=(a number here)'")
	fmt.Printf("\n\n")

	println("ls-watched\n")
	println("\t\tCommand to list all movies you have watched. To have a limited number listings  append '-limit=(a number here)'")
	fmt.Printf("\n\n")

	println(" add -title='title' -genre='genre' -year='year'\n")
	println("\t\t To add movie to the database")
	fmt.Printf("\n\n")

	println("rm -id='unique-id'\n")
	println("\t \t to remove a certain movie from the databse")
	println("\n\n")

	println("surprise-me -genre='genre'\n")
	println("\t\t Suggests a random movie from the database that you haven't watched. if genre is not specified it gives random movie.")
	fmt.Printf("\n\n")

}
