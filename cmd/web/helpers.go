package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"marcode.io/url-parser/pkg/models"
	"marcode.io/url-parser/pkg/parser"
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

func getLinkDetails(url string) models.LinkDetails {

	response, err := http.Get(url)

	if err != nil {
		log.Printf("error fetching url: %v", err)
	}

	responseBody, _ := io.ReadAll(response.Body)

	var details models.LinkDetails

	details.PageURL = url

	version, _ := parser.FindHTMLVersion(bytes.NewReader(responseBody))
	details.HTMLVersion = version

	title, _ := parser.FindTitle(bytes.NewReader(responseBody))
	details.Title = title

	headings, err := parser.FindAllHeadings(bytes.NewReader(responseBody))
	details.Headings = headings

	links, _ := parser.FindAllLinks(http.Get, bytes.NewReader(responseBody), url)
	details.Links = links

	hasLoginForm, _ := parser.CheckIfPageHasLoginForm(bytes.NewReader(responseBody))
	details.HasLoginForm = hasLoginForm

	return details
}
