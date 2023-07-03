package user

import "Creata21/snippetbox/internal/repository"

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo:repo,}
}

func (u *UserService) Insert() {

}

func (u *UserService) Authenticate() {
	
}

func (u *UserService) Get() {

}