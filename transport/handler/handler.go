package handler

import (
	"Creata21/snippetbox/internal/service"
	"Creata21/snippetbox/pkg/logger"
	"html/template"
)

type Handler struct {
	service  service.IService
	log      logger.Logger
	tmpl 	  map[string]*template.Template
}

func New(svc service.IService, l logger.Logger, template map[string]*template.Template) Handler {
	return Handler{service: svc, log: l, tmpl: template}
}
