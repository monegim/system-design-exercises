package app

import (
	"github.com/monegim/Reloader/inernal/pkg/cmd"
)

func Run() error {
	cmd := cmd.NewReloaderCommand()
	return cmd.Execute()
}
