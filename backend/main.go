package main

import (
	"flag"
	"os"

	"bitbucket.org/andreychernih/tweemote/commands"
	"bitbucket.org/andreychernih/tweemote/errors"
	"bitbucket.org/andreychernih/tweemote/models"
	"bitbucket.org/andreychernih/tweemote/rb"

	"github.com/golang/glog"
	"github.com/stvp/rollbar"

	"github.com/mitchellh/cli"
)

func init() {
	flag.Set("logtostderr", "false")
	flag.Set("stderrthreshold", "INFO")

	if os.Getenv("DEBUG") == "" {
		flag.Set("v", "1")
	} else {
		flag.Set("v", "2")
	}

	flag.Parse() // so that glog is happy
}

func main() {
	config, err := rb.LoadConfig("config.yaml")
	errors.Check(err)

	rollbar.Token = os.Getenv("ROLLBAR_TOKEN")
	rollbar.Environment = os.Getenv("ENV")
	defer rollbar.Wait()

	meta := commands.Meta{
		Config: config,
	}

	if os.Getenv("LOGDB") == "1" {
		models.Connect().LogMode(true)
	}

	c := cli.NewCLI("tweemote", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"impress": func() (cli.Command, error) {
			return &commands.ImpressCommand{
				Meta: meta,
			}, nil
		},
		"ingest": func() (cli.Command, error) {
			return &commands.IngestCommand{
				Meta: meta,
			}, nil
		},
		"persist": func() (cli.Command, error) {
			return &commands.PersistCommand{
				Meta: meta,
			}, nil
		},
		"process_user_queues": func() (cli.Command, error) {
			return &commands.ProcessUserQueuesCommand{
				Meta: meta,
			}, nil
		},
		"act": func() (cli.Command, error) {
			return &commands.ActCommand{
				Meta: meta,
			}, nil
		},
		"migrate": func() (cli.Command, error) {
			return &commands.MigrateCommand{
				Meta: meta,
			}, nil
		},
		"seed": func() (cli.Command, error) {
			return &commands.SeedCommand{
				Meta: meta,
			}, nil
		},
		"server": func() (cli.Command, error) {
			return &commands.ServerCommand{
				Meta: meta,
			}, nil
		},
		"stats": func() (cli.Command, error) {
			return &commands.StatsCommand{
				Meta: meta,
			}, nil
		},
	}
	glog.Infof("Starting %s", os.Args[1])
	exitStatus, err := c.Run()
	errors.Check(err)
	os.Exit(exitStatus)
}
