package handler

import (
	"Creata21/snippetbox/transport/middleware"
	"net/http"
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func Routes(h Handler) http.Handler {
	standardMiddleware := alice.New(middleware.RecoverPanic, middleware.LogRequest, middleware.SecureHeaders)
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(h.home))

	mux.Get("/snippet/create", http.HandlerFunc(h.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(h.createSnippet))

	mux.Get("/snippet/:id", http.HandlerFunc(h.showSnippet))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
