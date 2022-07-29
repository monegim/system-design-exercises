package main

import (
	"os"

	"github.com/monegim/Reloader/inernal/pkg/app"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}