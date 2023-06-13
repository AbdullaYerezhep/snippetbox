package server

import (
	"net/http"
	"Creata21/snippetbox/transport/middleware"
	"Creata21/snippetbox/transport/handler"
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func routes() http.Handler {
	standardMiddleware := alice.New(middleware.recoverPanic, middleware.logRequest, middleware.secureHeaders)
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.Home))

	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))

	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
