package main

import "log"

type ExampleServer struct {
	addr      string
	logger    *log.Logger
	isEnabled bool
}
