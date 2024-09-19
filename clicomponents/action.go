package clicomponents

import (
	"github.com/MahikaJaguste/todocli/tasks"
	"github.com/urfave/cli/v2"
)

func ActionFunc(cCtx *cli.Context) error {
	err := tasks.GetTasks()
	return err
}
