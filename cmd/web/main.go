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

	log.Println("Starting server on :8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)

}
