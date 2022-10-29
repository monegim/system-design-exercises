package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	conf := ReadFlags()
	stat, err := os.Stdin.Stat()
	if err != nil {
		logrus.WithError(err).Error("stdin error")
		os.Exit(1)
	}
	_ = conf
	fmt.Println(stat.Mode())
}
