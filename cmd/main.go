package main

import (
	"Creata21/snippetbox/app"
	"Creata21/snippetbox/config"
	"Creata21/snippetbox/pkg/logger"
	"flag"

	_ "github.com/lib/pq"
)


func main() {
	cfg := config.GetConfig()
	flag.StringVar(&cfg.Port, "addr", ":8080", "HTTP network address")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.SecretSession, "secret-session", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret session key")
	flag.Parse()

	log := logger.New()

	err := app.Run(cfg, log)
	if err != nil {
		log.ErrorLog.Fatal()
	}

}
