package parser

import (
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
)

func isInternalLink(href, baseURL string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}

	parentUri, err := url.Parse(baseURL)
	if err != nil {
		return false
	}

	if uri.Host != parentUri.Host {

		if strings.HasSuffix(uri.Host, parentUri.Host) {
			return true
		}

		return false
	}

	return true
}

func removeTrailingSlash(link string) string {
	return strings.TrimRight(link, "/")
}

func getUniqueLinks(links []Link) []Link {
	visited := make(map[string]bool)
	var uniqueLinks []Link

	for _, link := range links {
		if _, ok := visited[link.Href]; !ok {
			visited[link.Href] = true
			uniqueLinks = append(uniqueLinks, link)
		}
	}

	return uniqueLinks
}

func isAccessibleLink(url string) bool {
	response, err := http.Get(url)
	if err != nil {
		return false
	}

	if response.StatusCode == http.StatusOK {
		return true
	}

	return false
}

func getTextFromNode(node *html.Node) string {
	// Base Case
	if node.Type == html.TextNode {
		return node.Data
	}

	var text string
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		text += getTextFromNode(child)
	}

	return strings.Join(strings.Fields(text), " ")
}
