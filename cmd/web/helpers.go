package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type envelope map[string]interface{}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, templateData *templateData) {
	templateSet, ok := app.templatesCache[templateName]
	if !ok {
		app.serverErrorResponse(w, r, errors.New("error rendering template files"))
		return
	}

	err := templateSet.Execute(w, templateData)
	if err != nil {
		app.serverErrorResponse(w, r, errors.New("error rendering template files"))
		return
	}
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)

	return nil
}
