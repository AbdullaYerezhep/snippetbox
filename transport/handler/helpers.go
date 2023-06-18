package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (h *Handler) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	h.log.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (h *Handler) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (h *Handler) notFound(w http.ResponseWriter) {
	h.clientError(w, http.StatusNotFound)
}

func (h *Handler) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := h.tmpl[name]

	if !ok {
		h.serverError(w, fmt.Errorf("the template %s does not exists", name))
	}

	buf := new(bytes.Buffer)
	err := ts.Execute(buf, h.addDefaultData(td, r))

	if err != nil {
		h.serverError(w, err)
	}

	buf.WriteTo(w)
}

func (h *Handler) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	td.Flash = h.sessions.PopString(r, "flash")
	return td
}
