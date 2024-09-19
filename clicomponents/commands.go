package clicomponents

import (
	"github.com/MahikaJaguste/todocli/tasks"
	"github.com/urfave/cli/v2"
)

var Commands []*cli.Command = []*cli.Command{
	addCommand,
	showCommand,
}

var addCommand *cli.Command = &cli.Command{
	Name:      "add",
	Aliases:   []string{"a"},
	Usage:     "Add a pending task",
	Args:      true,
	ArgsUsage: "Enter task description, eg. add \"Finish assignment 1\"",
	Action: func(cCtx *cli.Context) error {
		taskDescription := cCtx.Args().First()
		createErr := tasks.CreateTask(taskDescription)
		return createErr
	},
}

var showCommand *cli.Command = &cli.Command{
	Name:    "show",
	Aliases: []string{"s"},
	Usage:   "Show all pending task",
	Args:    false,
	Action: func(cCtx *cli.Context) error {
		err := tasks.GetTasks()
		return err
	},
}
