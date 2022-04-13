package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home) // == mux.Handle("/", http.HandlerFunc(home)))
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	return standardMiddleware.Then(mux)

	//return logRequestd(app, secureHeaders(mux))
}
