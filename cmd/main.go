package main

import (
	"Creata21/snippetbox/app"
	"Creata21/snippetbox/config"
	"Creata21/snippetbox/pkg/logger"
	"flag"

	_ "github.com/lib/pq"
)

type Config struct {
	Addr      string
	StaticDir string
}

func main() {
	cfg := config.GetConfig()
	flag.StringVar(&cfg.Port, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	log := logger.New()

	err := app.Run(cfg, log)
	if err != nil {
		log.ErrorLog.Fatal()
	}

}
