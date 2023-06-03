package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr 		string
	StaticDir 	string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	cfg := new(Config)

	flag.StringVar(&cfg.Addr, "addr", ":8000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr: cfg.Addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Starting server at port %s!", cfg.Addr)

	err := srv.ListenAndServe()
		errorLog.Fatal(err)
}
