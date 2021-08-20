package main

import (
	"fmt"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.renderTemplate(w, r, "home.page.tmpl", nil)
}

func (app *application) showDetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here are the details about the submitted link")
}
