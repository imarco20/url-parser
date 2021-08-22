package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// envelope is used for composing different types of data (as key-value pairs)
// rendered into json and sent to response writer
type envelope map[string]interface{}

// renderTemplate searches for a template in the template cache and
// executes it passing the parameter template data
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

// writeJSON encodes data into JSON and writes it to the response writer,
// sets the content type header to application/json
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
