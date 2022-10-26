package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func Run(version string, args []string) {
	err := NewApp(version, time.Now()).Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewApp(version string, start time.Time) (app *cli.App) {
	app = cli.NewApp()
	app.Name = "gitlab-ci-pipelines-exporter"
	app.Version = version
	app.Usage = "Export metrics about Gitlab CI pipelines statuses"
	app.EnableBashCompletion = true
	app.Flags = cli.FlagsByName{
		&cli.StringFlag{
			Name:    "internal-monitoring-listener-address",
			Aliases: []string{"m"},
			EnvVars: []string{"GCPE_INTERNAL_MONITORING_LISTENER_ADDRESS"},
			Usage:   "internal monitoring listener address",
		},
	}
	app.Commands = cli.CommandsByName{

	}
	return
}
