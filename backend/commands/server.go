package commands

import (
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/server"
)

type ServerCommand struct {
	Meta
}

func (cmd ServerCommand) Help() string {
	return ""
}

func (cmd ServerCommand) Synopsis() string {
	return "starts API server"
}

func (cmd ServerCommand) Run(args []string) int {
	// Trigger connection so that we can turn on logging mode
	models.Connect()

	server.Start()
	return 0
}
