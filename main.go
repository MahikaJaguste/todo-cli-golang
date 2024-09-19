package main

import (
	"log"
	"os"

	"github.com/MahikaJaguste/todocli/clicomponents"
	"github.com/MahikaJaguste/todocli/db"
	"github.com/urfave/cli/v2"
)

func main() {

	initFileErr := db.InitFile()
	if initFileErr != nil {
		log.Fatalln("Failed to set up ToDo App database :(")
	}

	// App is the main structure of a cli application.
	app := &cli.App{
		Name:  "Todo List",
		Usage: "Add tasks, complete them and keep track of your life!",

		Commands: clicomponents.Commands,

		// executed when no valid subcommand is specified
		Action: clicomponents.ActionFunc,

		Flags: clicomponents.Flags,
	}

	// Run is the entry point to the cli app.
	// Parses the arguments slice and routes to the proper flag/args combination

	// os.Args Args hold the command-line arguments, starting with the program name.
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
