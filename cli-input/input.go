package cliinput

import (
	"flag"
	"os"
)

func HandleInput() {

	lsGenreCmd := flag.NewFlagSet("ls-genre", flag.ExitOnError)
	lsGenreCmdLim := lsGenreCmd.Uint("limit", 0, "Specify the maximum number of genre you want to see default 0 will list them all")

	surpriseMeCmd := flag.NewFlagSet("surprise-me", flag.ExitOnError)
	surpriseMeGenre := surpriseMeCmd.String("genre", "", "Specify the genre you want to watch a movie from. By default empty string will list them all.")

	addMovieCmd := flag.NewFlagSet("add", flag.ExitOnError)
	movieTitle := addMovieCmd.String("title", "", "Enter the movie title")
	movieGenre := addMovieCmd.String("genre", "", "Enter the movie genre")
	movieYear := addMovieCmd.Uint("year", 0, "Movie release year")

	lsWatchedCmd := flag.NewFlagSet("ls-watched", flag.ExitOnError)
	lsWatchedLim := lsWatchedCmd.Uint("limit", 0, "Specify the max number of watched movies you want to list. Default 0 will list them all")

	lsToWatchCmd := flag.NewFlagSet("ls-towatch", flag.ExitOnError)
	lsToWatchLim := lsToWatchCmd.Uint("limit", 0, "Specify the max number of unwatched movies you want to list. Default 0 will list them all")

	delCmd := flag.NewFlagSet("rm", flag.ExitOnError)
	delCmdId := delCmd.Uint("id", 0, "Delete a movie based on it's id")

	watchedCmd := flag.NewFlagSet("watched", flag.ExitOnError)
	watchedCmdId := watchedCmd.Uint("id", 0, "Enter the id of the movie to mark it watched")

	switch os.Args(1) {
	case "ls-genre":
		handleLsGenre(lsGenreCmdLim)
	case "surprise-me":
		handleSurpriseMe(surpriseMeGenre)
	case "add":
		handleAddMovie(movieTitle, movieGenre, movieYear)
	case "ls-watched":
		handleLsWatched(lsWatchedLim)
	case "ls-towatch":
		handleLsToWatch(lsToWatchLim)
	case "rm":
		handleRemoveById(delCmdId)
	case "watched":
		handleMoveToWatched(watchedCmdId)
	default:
		printHelp()
	}
}
