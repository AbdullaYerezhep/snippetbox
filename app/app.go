package app

import (
	"Creata21/snippetbox/config"
	"Creata21/snippetbox/internal/repository"
	"Creata21/snippetbox/internal/service"
	"Creata21/snippetbox/pkg/logger"
	"Creata21/snippetbox/pkg/postgres"
)

func Run(cfg config.Config, logger logger.Logger) error {
	

	db, err := postgres.OpenDB(cfg.DSN)

	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	defer db.Close()

	repo := repository.New(db, logger)
	service := service.New(repo, logger)
	return nil
}
