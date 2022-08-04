package main

import (
	"log"
	"net/http"
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

func (s *ExampleServer) ListenAndServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if s.isEnabled {
			s.logger.Println("responded with OK")
			w.WriteHeader(http.StatusOK)
		} else {
			s.logger.Println("responded with Error")
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/toggle", func(w http.ResponseWriter, r *http.Request) {
		s.isEnabled = !s.isEnabled
		s.logger.Println("toggled. Is enagled:", s.isEnabled)
		w.WriteHeader(http.StatusOK)
	})
	return http.ListenAndServe(s.addr, nil)
}
