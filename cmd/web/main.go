package main

import (
	"Creata21/snippetbox/app"
	"Creata21/snippetbox/config"
	"Creata21/snippetbox/pkg/logger"
	"Creata21/snippetbox/pkg/models/postgres"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Addr      string
	StaticDir string
}

type Application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	cfg := config.GetConfig()
	flag.StringVar(&cfg.Port, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	srv := &http.Server{
		Addr:     cfg.Port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	log := logger.New()

	err := app.Run(cfg, log)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
