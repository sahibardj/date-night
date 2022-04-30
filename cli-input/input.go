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

	switch os.Args[1] {
	case "ls-genre":
		lsGenreCmd.Parse(os.Args[2:])
		handleLsGenre(*lsGenreCmdLim)
	case "surprise-me":
		surpriseMeCmd.Parse(os.Args[2:])
		handleSurpriseMe(*surpriseMeGenre)
	case "add":
		addMovieCmd.Parse(os.Args[2:])
		handleAddMovie(*movieTitle, *movieGenre, *movieYear)
	case "ls-watched":
		lsWatchedCmd.Parse(os.Args[2:])
		handleLsWatched(*lsWatchedLim)
	case "ls-towatch":
		lsToWatchCmd.Parse(os.Args[2:])
		handleLsToWatch(*lsToWatchLim)
	case "rm":
		delCmd.Parse(os.Args[2:])
		handleRemoveById(*delCmdId)
	case "watched":
		watchedCmd.Parse(os.Args[2:])
		handleMovieWatched(*watchedCmdId)
	default:
		printHelp()
	}
}
