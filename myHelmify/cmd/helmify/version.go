package main

import "fmt"

var (
	version = "development"
	date    = "not set"
	commit  = "not set"
)

func printVersion() {
	fmt.Printf("Version:    %s\n", version)
	fmt.Printf("Build Time: %s\n", date)
	fmt.Printf("Git Commit: %s\n", commit)
}
