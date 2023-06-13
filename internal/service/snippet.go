package service

import "Creata21/snippetbox/pkg/models"

func (s service) Get(id int) (*models.Snippet, error) {
	return &models.Snippet{}, nil
}

func (s service) Insert(title, content string) (int, error) {
 return 0, nil
}

func (s service) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{}, nil
}