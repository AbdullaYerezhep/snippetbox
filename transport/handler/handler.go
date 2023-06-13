package handler

import (
	"Creata21/snippetbox/internal/service"
	"Creata21/snippetbox/pkg/logger"
)

type Handler struct {
	service service.IService
	log     logger.Logger
}

func New(svc service.IService, l logger.Logger) Handler {
	return Handler{service: svc, log: l}
}
