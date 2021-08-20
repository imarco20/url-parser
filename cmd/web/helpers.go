package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"marcode.io/url-parser/pkg/models"
	"marcode.io/url-parser/pkg/parser"
	"net/http"
)

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, templateData *templateData) {
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

func getLinkDetails(url string) models.LinkDetails {

	response, err := http.Get(url)

	if err != nil {
		log.Printf("error fetching url: %v", err)
	}

	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))

	var details models.LinkDetails

	details.PageURL = url

	version, _ := parser.FindHTMLVersion(bytes.NewReader(responseBody))
	details.HTMLVersion = version

	title, _ := parser.FindTitle(bytes.NewReader(responseBody))
	details.Title = title

	headings, err := parser.FindAllHeadings(bytes.NewReader(responseBody))
	details.Headings = headings

	links, _ := parser.FindAllLinks(bytes.NewReader(responseBody), url)
	details.Links = links

	hasLoginForm, _ := parser.CheckIfPageHasLoginForm(bytes.NewReader(responseBody))
	details.HasLoginForm = hasLoginForm

	return details
}
