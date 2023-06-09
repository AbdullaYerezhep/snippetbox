package main

import (
	"Creata21/snippetbox/config"
	"Creata21/snippetbox/pkg/models/postgres"
	"context"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
	Addr      string
	StaticDir string
}

type application struct {
	snippets *postgres.SnippetModel
	errorLog *log.Logger
	infoLog  *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	cfg := config.GetConfig()

	flag.StringVar(&cfg.Port, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()


	log.Println("Configs are parsed")
	log.Println(cfg.DSN)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.DSN)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html/")

	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgres.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     cfg.Port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server at port %s!", cfg.Port)

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
