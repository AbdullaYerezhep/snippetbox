package repository

import (
	"Creata21/snippetbox/pkg/models"
	"database/sql"
)

type IDb interface {
	Insert(title, content string) (int, error)
	Get(id int) (*models.Snippet, error)
	Latest() ([]*models.Snippet, error) 
}

type repository struct {
	DB *sql.DB
}

func New(db *sql.DB) IDb {
	return repository{DB: db}
}
