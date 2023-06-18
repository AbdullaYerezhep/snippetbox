package handler

import (
	"Creata21/snippetbox/internal/service"
	"Creata21/snippetbox/pkg/logger"
	"html/template"

	"github.com/golangcollege/sessions"
)

type Handler struct {
	service  service.IService
	log      logger.Logger
	tmpl 	  map[string]*template.Template
	sessions  sessions.Session
}

func New(svc service.IService, l logger.Logger, template map[string]*template.Template, session *sessions.Session) *Handler {
	return &Handler{service: svc, log: l, tmpl: template, sessions: *session}
}
