package service

import (
	"Creata21/snippetbox/pkg/models"
	"strings"
	"unicode/utf8"
)

func (s service) Get(id int64) (*models.Snippet, error) {
	res, err := s.repository.Get(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s service) Insert(title, content string) (int, map[string]string) {
	errors := make(map[string]string)
	if strings.TrimSpace(title) == "" {
		errors["title"] = "The title field cannot be empty"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This title field is too long (maximum is 100 characters)"
	}

	if strings.TrimSpace(content) == "" {
		errors["content"] = "The title content cannot be empty"
	}

	if len(errors) > 0 {
		return 0, errors
	}
	res, err := s.repository.Insert(title, content)

	errors["database"] = err.Error()
	if len(errors) > 0 {
		return 0, errors
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
