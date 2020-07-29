package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// create middleware chain
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// create dynamic middleware chain
	dynamicMiddleware := alice.New(app.session.Enable)

	mux := pat.New()

	// routes for snippets
	mux.Get("/", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.home)))
	mux.Get("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(http.HandlerFunc(app.createSnippetForm)))
	mux.Post("/snippet/create", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(http.HandlerFunc(app.createSnippet)))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.showSnippet)))

	// routes for user
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.userSignupForm)))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.userSignup)))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.userLoginForm)))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(http.HandlerFunc(app.userLogin)))
	mux.Post("/user/logout", dynamicMiddleware.Append(app.requireAuthenticatedUser).ThenFunc(http.HandlerFunc(app.userLogout)))

	// File server to serve static files
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Register filesServer as handler
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)

}
