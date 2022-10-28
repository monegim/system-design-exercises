package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		logrus.WithError(err).Error("stdin error")
		os.Exit(1)
	}
	_ = stat
}
