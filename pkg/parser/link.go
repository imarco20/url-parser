package parser

import (
	"golang.org/x/net/html"
	"io"
)

type Link struct {
	Href string
}

type LinkCount struct {
	Internal     int
	External     int
	InAccessible int
}

func FindAllLinks(body io.Reader, pageURL string) (LinkCount, error) {
	document, err := html.Parse(body)
	if err != nil {
		return LinkCount{}, err
	}

	var links []Link
	linkNodes := getNodes(document, "a")

	for _, node := range linkNodes {
		links = append(links, node.buildLink())
	}

	uniqueLinks := getUniqueLinks(links)

	var linkCount LinkCount
	for _, link := range uniqueLinks {
		if isInternalLink(link.Href, pageURL) {
			linkCount.Internal++
		} else {
			linkCount.External++
		}

		if !isAccessibleLink(link.Href) {
			linkCount.InAccessible++
		}
	}

	return linkCount, nil
}
