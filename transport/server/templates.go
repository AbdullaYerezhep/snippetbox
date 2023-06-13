package server

import (
	"Creata21/snippetbox/pkg/models"
	"html/template"
	"log"
	"path/filepath"
	"time"
)

type TemplateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
	CurrentYear int
}

func humanDate(t time.Time) string {
	return t.Format("01 Jan 2006 at 15:04")
	
}

var functions = template.FuncMap{
	"humanDate" : humanDate,
}

func NewTemplateCache(dir string) (map[string]*template.Template, error) {
	log.Println(dir)
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*page.tmpl"))

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}	
		
		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.tmpl"))	

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*partial.tmpl"))

		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}