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

func main() {
	cfg := new(Config)

	flag.StringVar(&cfg.Addr, "addr", ":8000", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr: cfg.Addr,
		ErrorLog: errLog,
		Handler: mux,
	}

	infoLog.Printf("Starting server at port %s!", cfg.Addr)

	err := srv.ListenAndServe()
		errLog.Fatal(err)
}
