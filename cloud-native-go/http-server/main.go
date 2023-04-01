package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",getRoot)
	http.HandleFunc("/hello",getHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request)  {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request)  {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello Http\n")
}