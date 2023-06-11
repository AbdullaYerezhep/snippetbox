package main

import (
	"Creata21/snippetbox/pkg/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request)  {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Snippets: snippets}

	app.render(w, r, "home.page.tmpl", data)

}

func (app *Application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	data := &templateData{Snippet: s}

	app.render(w, r, "show.page.tmpl", data)
}

func (app *Application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", nil)
}

func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	errors := make(map[string]string)

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	if strings.TrimSpace(title) == "" {
		errors["titile"] = "The title field cannot be empty"
	}else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This title field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "The title content cannot be empty"
	}

	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	id, err := app.snippets.Insert(title, content)
	fmt.Println(id)
	if err != nil {
		
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
