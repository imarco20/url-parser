package main

import "net/http"

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(app.homeHandler))
	router.Handle("/details", http.HandlerFunc(app.showDetailsHandler))
	router.Handle("/healthcheck", http.HandlerFunc(app.healthCheckHandler))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	return router
}
