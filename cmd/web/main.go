package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on http://localhost:4000")
	log.Println("Starting server on http://localhost:4000/snippet/view")
	log.Println("Starting server on http://localhost:4000/snippet/create")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
