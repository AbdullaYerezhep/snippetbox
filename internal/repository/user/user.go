package user

import (
	"Creata21/snippetbox/pkg/models"
	"database/sql"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db:db,
	}
}

func (u *UserStorage) Insert(name, email, password string) error{
	return nil
}

func (u * UserStorage) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (u *UserStorage) Get(id int) (*models.User, error) {
	return &models.User{}, nil
}