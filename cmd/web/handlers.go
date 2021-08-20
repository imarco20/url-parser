package main

import (
	"fmt"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func (app *application) showDetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here are the details about the submitted link")
}
