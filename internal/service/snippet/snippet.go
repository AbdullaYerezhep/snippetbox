package snippet

import (
	"Creata21/snippetbox/internal/repository"
	"Creata21/snippetbox/pkg/models"
)

type SnippetService struct {
	repo *repository.Repository
}

func NewSnippetService(repo *repository.Repository) *SnippetService {
	return &SnippetService{
		repo: repo,
	}
}

func (s *SnippetService) Get(id int64) (*models.Snippet, error) {
	res, err := s.repo.SnippetRepo.Get(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SnippetService) Insert(title, content string) (int, error) {

	res, err := s.repo.SnippetRepo.Insert(title, content)

	if err != nil {
		return 0, err
	}

	return res, nil
}

func (s *SnippetService) Latest() ([]*models.Snippet, error) {
	res, err := s.repo.SnippetRepo.Latest()
	if err != nil {
		return nil, err
	}
	return res, nil
}
