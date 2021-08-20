package main

import (
	"marcode.io/url-parser/pkg/forms"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		app.renderTemplate(w, r, "home.page.tmpl", &templateData{Form: forms.New(nil)})

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *application) showDetailsHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:

		err := r.ParseForm()
		if err != nil {
			http.Error(w, "invalid form input", http.StatusBadRequest)
			return
		}

		form := forms.New(r.PostForm)
		form.Required("link")
		form.MatchesPattern("link", forms.URLRegExp)

		if !form.Valid() {
			app.renderTemplate(w, r, "home.page.tmpl", &templateData{Form: form})
		} else {
			details := getLinkDetails(form.Values.Get("link"))
			app.renderTemplate(w, r, "details.page.tmpl", &templateData{Link: details})
		}

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
