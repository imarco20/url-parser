package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

type config struct {
	port int
}

type application struct {
	config
	logger         *log.Logger
	templatesCache map[string]*template.Template
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	templatesCache, err := cacheAllTemplates("./ui/html/")
	if err != nil {
		logger.Println(err)
	}

	app := &application{
		config:         cfg,
		logger:         logger,
		templatesCache: templatesCache,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	logger.Printf("starting server on %s", server.Addr)
	err = server.ListenAndServe()
	logger.Fatal(err)
}
