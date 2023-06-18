package repository

import (
	"Creata21/snippetbox/pkg/models"
	"Creata21/snippetbox/pkg/logger"
	"database/sql"
)

type IDb interface {
	Insert(title, content string) (int, error)
	Get(id int64) (*models.Snippet, error)
	Latest() ([]*models.Snippet, error)
}

type repository struct {
	DB  *sql.DB
	log logger.Logger
}

func New(db *sql.DB, l logger.Logger) IDb {
	return &repository{DB: db, log: l}
}
