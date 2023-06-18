package service

import (
	"Creata21/snippetbox/internal/repository"
	"Creata21/snippetbox/pkg/logger"
	"Creata21/snippetbox/pkg/models"
)

type IService interface {
	Insert(title, content string) (int, error)
	Get(id int64) (*models.Snippet, error)
	Latest() ([]*models.Snippet, error)
}

type service struct {
	repository repository.IDb
	log        logger.Logger
}

func New(r repository.IDb, l logger.Logger) IService {
	return &service{repository: r, log: l}
}
