package main

import (
	"log"
	"os"
)

type ExampleServer struct {
	addr      string
	logger    *log.Logger
	isEnabled bool
}

func NewExampleServer(addr string) *ExampleServer {
	return &ExampleServer{
		addr: addr,
		logger: log.New(os.Stdout, "Server\t", log.LstdFlags),
		isEnabled: true,
	}
}
