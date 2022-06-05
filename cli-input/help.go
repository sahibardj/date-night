package cliinput

import "fmt"

func printHelp() {
	println("WELCOME TO DATE NIGHT MOVIE SELECTOR")

	fmt.Printf("\n\n")

	println("To list out all the genre available type 'ls-genre'")
	println("To have a limited number of genre listed append '-limit=(a number here)'")
	fmt.Printf("\n\n")

	println("To list out all the movies you havn't watched 'ls-towatch'")
	println("To have a limited number of genre listed append '-limit=(a number here)'")
	fmt.Printf("\n\n")

	println("To list out all the movies you have watched 'ls-watched'")
	println("To have a limited number of genre listed append '-limit=(a number here)'")
	fmt.Printf("\n\n")

	println("To add a movie")
	println(" add -title='title' -genre='genre' -year='year'")
	fmt.Printf("\n\n")

	println("SELECT A RANDOM MOVIE")
	println(("Surprise-me -genre='genre'")
	fmt.Printf("\n\n")

	println("Watched a movie and don't want it on your list?\n Or just don't want a movie on your list anymore? Remove them simply\t ' rm -id[unique id associated] ' \n\n")
	fmt.Printf("\n\n")

}
