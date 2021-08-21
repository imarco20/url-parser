package parser

import (
	"golang.org/x/net/html"
	"io"
)

type Title struct {
	Value string
}

func FindTitle(body io.Reader) (string, error) {
	document, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	nodes := getNodes(document, "title")
	if len(nodes) == 0 {
		return "", ErrPageTitleNotFound
	}

	titleNode := nodes[0]

	var title Title
	title = titleNode.buildTitle()

	return title.Value, nil
}
