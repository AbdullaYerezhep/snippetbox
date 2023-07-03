package handler

import (
	"Creata21/snippetbox/pkg/forms"
	"Creata21/snippetbox/pkg/models"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := h.service.SnippetService.Latest()
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

	s, err := h.service.SnippetService.Get(int64(id))

	if err == models.ErrNoRecord {
		h.notFound(w)
	} else if err != nil {
		h.serverError(w, err)
		return
	}
	data := &templateData{
		Snippet: s,
	}
	h.render(w, r, "show.page.tmpl", data)
}

func (h *Handler) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	h.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (h *Handler) createSnippet(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title", "content")
	form.MaxLength("title", 100)

	if !form.Valid() {
		h.render(w, r, "create.page.tmpl", &templateData{
			Form: form,
		})
		return
	}

	id, err := h.service.SnippetService.Insert(form.Get("title"), form.Get("content"))

	if err != nil{
		h.serverError(w, err)
		return
	}

	h.sessions.Put(r, "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
