package main

import (
	"marcode.io/url-parser/pkg/forms"
	"net/http"
)

// homeHandler handles GET requests to the home page and renders its template
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFoundResponse(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.renderTemplate(w, r, "home.page.tmpl", &templateData{Form: forms.New(nil)})

	default:
		app.methodNotAllowedResponse(w, r)
	}
}

// showDetailsHandler handles POST requests sent by submitting the form in the home page
// and displays URL details
func (app *application) showDetailsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		err := r.ParseForm()
		if err != nil {
			app.badRequestResponse(w, r, err)
			return
		}

		form := forms.New(r.PostForm)
		form.Required("link")
		form.MatchesPattern("link", forms.URLRegExp)

		if !form.Valid() {
			app.renderTemplate(w, r, "home.page.tmpl", &templateData{Form: form})
		} else {
			details := app.parser(form.Values.Get("link"))
			app.renderTemplate(w, r, "details.page.tmpl", &templateData{Link: details})
		}

	default:
		app.methodNotAllowedResponse(w, r)
	}
}

// healthCheckHandler handles requests to check the application is up and running
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.methodNotAllowedResponse(w, r)
		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"health": "the application is working properly"})
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
