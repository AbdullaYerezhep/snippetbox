package main

import (
	"Creata21/snippetbox/pkg/models/mysql"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Addr 		string
	StaticDir 	string
}

type application struct {
	snippets *mysql.SnippetModel
	errorLog *log.Logger
	infoLog  *log.Logger
}
const (
    username = "root"
    password = "AYcreata21$"
    host     = "172.17.0.4"
    port     = "3306"
    database = "snippetbox"
)

func main() {
	cfg := new(Config)
	
	flag.StringVar(&cfg.Addr, "addr", ":8000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()
	
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
	db, err  := openDB(dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &application {
		errorLog: errorLog,
		infoLog: infoLog,
		snippets: &mysql.SnippetModel{DB: db},
	}

	srv := &http.Server {
		Addr: cfg.Addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server at port %s!", cfg.Addr)

	err = srv.ListenAndServe()
		errorLog.Fatal(err)
}


func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}