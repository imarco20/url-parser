package parser

import (
	"golang.org/x/net/html"
	"io"
)

// HTMLTag represents a HTML tag
type HTMLTag struct {
	Version string
}

// FindHTMLVersion finds the version of a HTML document
func FindHTMLVersion(body io.Reader) (string, error) {
	document, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	nodes := getNodes(document, "html")

	htmlNode := nodes[0]

	var htmlTag HTMLTag
	htmlTag = htmlNode.buildVersion()

	if htmlTag.Version == "" {
		return "", ErrHTMLVersionNotFound
	}

	return htmlTag.Version, nil
}
