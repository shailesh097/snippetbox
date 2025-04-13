package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a specific snippet."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on: http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
