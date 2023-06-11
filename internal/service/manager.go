package service

import "Creata21/snippetbox/internal/repository"

type IService interface {
	Create() 
	Insert()
	Latest()
}

type service struct {
	repository repository.IDb
}

func New(r repository.IDb) IService {
	return service{repository: r}
}