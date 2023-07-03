package service

import (
	"Creata21/snippetbox/internal/repository"
	"Creata21/snippetbox/internal/service/snippet"
	"Creata21/snippetbox/internal/service/user"
	"Creata21/snippetbox/pkg/logger"
)

type Service struct {
	UserService 	*user.UserService
	SnippetService *snippet.SnippetService
	log        logger.Logger
}

func New(r *repository.Repository, l logger.Logger) *Service {
	return &Service{UserService: user.NewUserService(r), SnippetService: snippet.NewSnippetService(r), log: l}
}
