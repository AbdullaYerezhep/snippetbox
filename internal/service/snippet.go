package service

import "Creata21/snippetbox/pkg/models"

func (s service) Get(id int64) (*models.Snippet, error) {
	res, err := s.repository.Get(id)

	if err != nil {
		return nil, err
	}


	return res, nil
}

func (s service) Insert(title, content string) (int, error) {
	res, err := s.repository.Insert(title, content)

	if err != nil {
		return 0, err
	}

 	return res, nil
}

func (s service) Latest() ([]*models.Snippet, error) {
	res, err := s.repository.Latest()
	if err != nil {
		return nil, err
	}
	return res, nil
}