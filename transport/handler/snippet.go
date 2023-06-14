package handler

import (
	"Creata21/snippetbox/pkg/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := h.service.Latest()
	if err != nil {
		h.serverError(w, err)
		return
	}
	data := &templateData{Snippets: snippets}

	h.render(w, r, "home.page.tmpl", data)

}

func (h *Handler) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		h.notFound(w)
		return
	}
	s, err := h.service.Get(int64(id))
	if err == models.ErrNoRecord {
		h.notFound(w)
	} else if err != nil {
		h.serverError(w, err)
		return
	}
	data := &templateData{Snippet: s}

	h.render(w, r, "show.page.tmpl", data)
}

func (h *Handler) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "create.page.tmpl", nil)
}

func (h *Handler) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	errors := make(map[string]string)

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	if strings.TrimSpace(title) == "" {
		errors["titile"] = "The title field cannot be empty"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This title field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "The title content cannot be empty"
	}

	if len(errors) > 0 {
		fmt.Fprint(w, errors)
		return
	}

	id, err := h.service.Insert(title, content)
	fmt.Println(id)
	if err != nil {

		h.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
