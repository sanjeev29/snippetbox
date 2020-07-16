package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/snippet", showSnippet)

	// File server to serve static files
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Register filesServer as handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)

}
