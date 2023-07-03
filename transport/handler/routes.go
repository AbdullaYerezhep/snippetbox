package handler

import (
	"Creata21/snippetbox/transport/middleware"
	"net/http"
	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func Routes(h Handler) http.Handler {
	standardMiddleware := alice.New(middleware.RecoverPanic, middleware.LogRequest, middleware.SecureHeaders)
	dynamicMiddleware := alice.New(h.sessions.Enable)
	mux := pat.New()

	mux.Get("/", dynamicMiddleware.ThenFunc(h.home))

	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(h.createSnippetForm))
	mux.Post("/snippet/create",dynamicMiddleware.ThenFunc(h.createSnippet))

	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(h.showSnippet))

	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(h.signupUserForm))
	mux.Post("/user/sigup", dynamicMiddleware.ThenFunc(h.signupUser))

	mux.Get("/user/login", dynamicMiddleware.ThenFunc(h.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(h.loginUser))

	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(h.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
