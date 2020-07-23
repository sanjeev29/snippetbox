package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// create middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.HandleFunc("/snippet", app.showSnippet)

	// File server to serve static files
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Register filesServer as handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)

}
