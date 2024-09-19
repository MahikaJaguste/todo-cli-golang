package clicomponents

import "github.com/urfave/cli/v2"

var Flags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:  "date",
		Value: "english",
		Usage: "language for greeting",
		// Destination: &language,
	},
	&cli.BoolFlag{
		Name:    "foo",
		Usage:   "foo greeting",
		Aliases: []string{"f"},
		// Count:   &count,
	},
}
