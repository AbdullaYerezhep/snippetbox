package handler

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

func (h Handler) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s",err.Error(), debug.Stack())
	h.log.errorLog.Output(2, trace)
	
	http.Error(w, http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}

func  clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func  notFound(w http.ResponseWriter) {
	clientError(w, http.StatusNotFound)
}

func  render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]

	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exists", name))
	}

	buf := new(bytes.Buffer)
	err := ts.Execute(buf, app.addDefaultData(td, r))

	if err != nil {
		app.serverError(w, err)
	}

	buf.WriteTo(w)
}

func (app *Application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td
}