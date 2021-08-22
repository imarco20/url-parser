package parser

import (
	"golang.org/x/net/html"
	"io"
)

// HeadingCount includes the count of each HTML heading level
type HeadingCount struct {
	HOne   int
	HTwo   int
	HThree int
	HFour  int
	HFive  int
	HSix   int
}

// FindAllHeadings finds all HTML nodes of every heading level
// and stores these counts in a HeadingCount object
func FindAllHeadings(body io.Reader) (HeadingCount, error) {
	document, err := html.Parse(body)
	if err != nil {
		return HeadingCount{}, err
	}

	var count HeadingCount

	count.HOne = len(getNodes(document, "h1"))

	count.HTwo = len(getNodes(document, "h2"))

	count.HThree = len(getNodes(document, "h3"))

	count.HFour = len(getNodes(document, "h4"))

	count.HFive = len(getNodes(document, "h5"))

	count.HSix = len(getNodes(document, "h6"))

	return count, nil
}
