package main

import (
	"gitlab-ci-pipelines-exporter/internal/cli"
	"os"
)

var version = "devel"

func main() {
	cli.Run(version, os.Args)
}
