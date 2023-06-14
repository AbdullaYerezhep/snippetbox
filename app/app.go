package app

import (
	"Creata21/snippetbox/config"
	"Creata21/snippetbox/internal/repository"
	"Creata21/snippetbox/internal/service"
	"Creata21/snippetbox/pkg/logger"
	"Creata21/snippetbox/pkg/postgres"
	handler "Creata21/snippetbox/transport/handler"
	"net/http"
)

func Run(cfg *config.Config, logger logger.Logger) error {

	db, err := postgres.OpenDB(cfg.DSN)

	if err != nil {
		logger.ErrorLog.Fatal(err)
	}

	defer db.Close()

	repo := repository.New(db, logger)

	service := service.New(repo, logger)

	template, err := handler.NewTemplateCache("./ui/html")
	
	if err != nil {
		logger.ErrorLog.Fatal("Could not cache template!")
	}

	handle := handler.New(service, logger, template)
	srv := &http.Server{
		Addr:     cfg.Port,
		ErrorLog: logger.ErrorLog,
		Handler:  handler.Routes(handle),
	}

	srv.ListenAndServe()

	return nil
}
