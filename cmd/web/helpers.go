package main

import "net/http"

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, templateData interface{}) {
	templateSet, ok := app.templatesCache[templateName]
	if !ok {
		http.Error(w, "error rendering template files", http.StatusInternalServerError)
		return
	}

	err := templateSet.Execute(w, templateData)
	if err != nil {
		http.Error(w, "error rendering template files", http.StatusInternalServerError)
		return
	}
}
