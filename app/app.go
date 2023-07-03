package app

import (
	"Creata21/snippetbox/config"
	"Creata21/snippetbox/internal/repository"
	"Creata21/snippetbox/internal/service"
	"Creata21/snippetbox/pkg/logger"
	"Creata21/snippetbox/pkg/postgres"
	handler "Creata21/snippetbox/transport/handler"
	"net/http"
	"time"

	"github.com/golangcollege/sessions"
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

	session := sessions.New([]byte(cfg.SecretSession))
	session.Lifetime = 12 * time.Hour

	handle := handler.New(service, logger, template, session)
	
	srv := &http.Server{
		Addr:     cfg.Port,
		ErrorLog: logger.ErrorLog,
		Handler:  handler.Routes(*handle),
		IdleTimeout: time.Minute,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.InfoLog.Printf("Server is running on port %s", srv.Addr)

	err = srv.ListenAndServe()

	logger.ErrorLog.Fatalln(err)
	
	return nil
}
