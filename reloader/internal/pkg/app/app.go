package app

import (
	"github.com/monegim/Reloader/internal/pkg/cmd"
)

func Run() error {
	cmd := cmd.NewReloaderCommand()
	return cmd.Execute()
}
