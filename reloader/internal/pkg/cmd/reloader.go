package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

func NewReloaderCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reloader",
		Short:   "A watcher for your Kubernetes cluster",
		PreRunE: validateFlags,
		Run:     startReloader,
	}
	return cmd
}

func validateFlags(*cobra.Command, []string) error {
	err := fmt.Sprintf(("must be one of: "))
	return errors.New(err)
}

func startReloader(cmd *cobra.Command, args []string) {}
