package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"marcode.io/url-parser/pkg/models"
	"net/http"
	"os"
	"os/signal"
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
	parser         models.Parser
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
		parser:         models.GetLinkDetails,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	logger.Printf("starting server on %s", server.Addr)

	go func() {
		err = server.ListenAndServe()
		logger.Fatal(err)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	// Graceful Shutdown for the server
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = server.Shutdown(tc)
	if err != nil {
		logger.Println(err)
	}
}
