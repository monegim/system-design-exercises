package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",getRoot)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request)  {
	
}