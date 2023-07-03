package repository

import (
	"Creata21/snippetbox/internal/repository/snippet"
	"Creata21/snippetbox/internal/repository/user"
	"Creata21/snippetbox/pkg/logger"
	"database/sql"
)

type Repository struct {
	UserRepo    *user.UserStorage
	SnippetRepo *snippet.SnippetStorage
	log         logger.Logger
}

func NewRepository(db *sql.DB, l logger.Logger) *Repository {
	return &Repository{UserRepo: user.NewUserStorage(db), SnippetRepo: snippet.NewSnippetStorage(db), log: l}
}
