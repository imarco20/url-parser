package parser

import (
	"golang.org/x/net/html"
	"io"
)

type HTMLTag struct {
	Version string
}

func FindHTMLVersion(body io.Reader) (string, error) {
	// TODO: Extract in a separate function
	document, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	nodes := getNodes(document, "html")

	// TODO: Handle Error in a user friendly way
	if len(nodes) == 0 {
		return "", err
	}

	htmlNode := nodes[0]

	var htmlTag HTMLTag
	htmlTag = htmlNode.buildVersion()

	return htmlTag.Version, nil
}
