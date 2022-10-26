package cmd

import (
	"time"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var start time.Time

func exit(exitCode int, err error) cli.ExitCoder{
	defer log.WithFields(
		log.Fields{
			"execution-time": time.Since(start),
		},
	).Debug("exited..")
	if err != nil {
		log.WithError(err).Error()
	}
	return cli.Exit("", exitCode)
}
func ExecWrapper(f func (ctx *cli.Context)(int, error)) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return exit(f(ctx))
	}
}