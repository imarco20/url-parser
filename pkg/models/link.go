package models

import (
	"bytes"
	"io"
	"log"
	"marcode.io/url-parser/pkg/parser"
	"net/http"
)

// LinkDetails contains all the required details about a web page link
type LinkDetails struct {
	PageURL      string
	HTMLVersion  string
	Title        string
	Headings     parser.HeadingCount
	Links        parser.LinkCount
	HasLoginForm bool
}

// Parser is a type for abstracting the signature of a
// function that takes in a url and returns a LinkDetails object
// It's used to mock the behavior of GetLinkDetails in tests
type Parser func(url string) LinkDetails

// GetLinkDetails returns all the details of the parameter URL
func GetLinkDetails(url string) LinkDetails {

	response, err := http.Get(url)

	if err != nil {
		log.Printf("error fetching url: %v", err)
	}

	responseBody, _ := io.ReadAll(response.Body)

	var details LinkDetails

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
