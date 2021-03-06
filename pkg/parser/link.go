package parser

import (
	"golang.org/x/net/html"
	"io"
)

// Link is a type for representing HTML anchor tags
type Link struct {
	Href string
}

// LinkCount includes the count of internal, external, and in-accessible links in a web page
type LinkCount struct {
	Internal     int
	External     int
	InAccessible int
}

// FindAllLinks finds all HTML anchor nodes and stores
// their details in a LinkCount object
func FindAllLinks(hg HttpGetter, body io.Reader, pageURL string) (LinkCount, error) {
	document, err := html.Parse(body)
	if err != nil {
		return LinkCount{}, err
	}

	var links []Link
	linkNodes := getNodes(document, "a")

	for _, node := range linkNodes {
		links = append(links, node.buildLink(pageURL))
	}

	uniqueLinks := getUniqueLinks(links)
	accessibleChannel := make(chan struct {
		string
		bool
	})

	var linkCount LinkCount
	for _, link := range uniqueLinks {
		if isInternalLink(link.Href, pageURL) {
			linkCount.Internal++
		} else {
			linkCount.External++
		}

		go func(url string) {
			accessibleChannel <- struct {
				string
				bool
			}{url, isAccessibleLink(hg, url)}
		}(link.Href)
	}

	for i := 0; i < len(uniqueLinks); i++ {
		result := <-accessibleChannel
		if !result.bool {
			linkCount.InAccessible++
		}
	}

	return linkCount, nil
}
