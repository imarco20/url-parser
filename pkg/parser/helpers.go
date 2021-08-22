package parser

import (
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"strings"
)

// HttpGetter is a type for functions that send HTTP GET requests
// to the specified URL. This is used to wrap sending HTTP requests
type HttpGetter func(string) (*http.Response, error)

// isInternalLink compares a link to it's parent URL
// to check if they belong to the same host
func isInternalLink(href, baseURL string) bool {
	childHost, err := getUrlHost(href)
	if err != nil {
		return false
	}

	parentHost, err := getUrlHost(baseURL)
	if err != nil {
		return false
	}

	if childHost != parentHost {
		// Cases 1 and 2: ChildURL is a sub path of the Parent URL
		// 1. The ParentURL Host doesn't start with www
		// 2. The ParentURL Host starts with www
		// Case 3: ChildURL is a relative path
		if strings.HasSuffix(childHost, parentHost) || strings.HasSuffix(childHost, parentHost[3:]) || childHost == "" {

			return true
		}

		return false
	}

	return true
}

// removeTrailingSlash removes a trailing slash from a URL
func removeTrailingSlash(link string) string {
	return strings.TrimRight(link, "/")
}

// getUniqueLinks filters out links that appear more than once in the page
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

// isAccessibleLink sends a GET request to the URL parameter and checks
// whether the response has status 200 OK
func isAccessibleLink(hg HttpGetter, url string) bool {
	response, err := hg(url)
	if err != nil {
		return false
	}

	if response.StatusCode == http.StatusOK {
		return true
	}

	return false
}

// getTextFromNode returns the text of a HTML node
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

// getUrlHost returns the host URI of the input link
func getUrlHost(href string) (string, error) {
	uri, err := url.Parse(href)
	if err != nil {
		return "", err
	}

	return uri.Host, nil
}
