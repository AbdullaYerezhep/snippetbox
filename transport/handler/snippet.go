package handler

import (
	"Creata21/snippetbox/pkg/models"
	"fmt"
	"net/http"
	"strconv"
	
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

	// errors := make(map[string]string)

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	id, errors := h.service.Insert(title, content)

	if len(errors) > 0 {
		h.render(w, r, "create.page.tmpl", &templateData{
			FormErrors:errors,
			FormData: r.PostForm,
		})
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
