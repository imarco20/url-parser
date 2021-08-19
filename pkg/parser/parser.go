package parser

import (
	"golang.org/x/net/html"
	"io"
	"net/url"
	"strings"
)

type Title struct {
	Value string
}

type HeadingCount struct {
	HOne   int
	HTwo   int
	HThree int
	HFour  int
	HFive  int
	HSix   int
}

type Link struct {
	Href string
}

type LinkCount struct {
	Internal int
	External int
}

func FindTitle(body io.Reader) (string, error) {
	document, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	nodes := getNodes(document, "title")
	if len(nodes) == 0 {
		return "", err
	}

	var title Title
	title = buildTitle(nodes[0])

	return title.Value, nil
}

func FindAllHeadings(body io.Reader) (HeadingCount, error) {
	document, err := html.Parse(body)
	if err != nil {
		return HeadingCount{}, err
	}

	var count HeadingCount

	headingOneNodes := getNodes(document, "h1")
	count.HOne = len(headingOneNodes)

	headingTwoNodes := getNodes(document, "h2")
	count.HTwo = len(headingTwoNodes)

	headingThreeNodes := getNodes(document, "h3")
	count.HThree = len(headingThreeNodes)

	headingFourNodes := getNodes(document, "h4")
	count.HFour = len(headingFourNodes)

	headingFiveNodes := getNodes(document, "h5")
	count.HFive = len(headingFiveNodes)

	headingSixNodes := getNodes(document, "h6")
	count.HSix = len(headingSixNodes)

	return count, nil
}

func FindAllLinks(body io.Reader, pageURL string) (LinkCount, error) {
	document, err := html.Parse(body)
	if err != nil {
		return LinkCount{}, err
	}

	var links []Link
	linkNodes := getNodes(document, "a")

	for _, node := range linkNodes {
		links = append(links, buildLink(node))
	}

	uniqueLinks := getUniqueLinks(links)

	var linkCount LinkCount
	for _, link := range uniqueLinks {
		if isInternalLink(link.Href, pageURL) {
			linkCount.Internal++
		} else {
			linkCount.External++
		}
	}

	return linkCount, nil
}

func getNodes(node *html.Node, nodeType string) []*html.Node {

	// Base Case
	if node.Type == html.ElementNode && node.Data == nodeType {
		return []*html.Node{node}
	}

	var nodes []*html.Node
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodes = append(nodes, getNodes(child, nodeType)...)
	}

	return nodes
}

func buildTitle(node *html.Node) (title Title) {
	title.Value = getTextFromNode(node)
	return
}

func buildLink(node *html.Node) (link Link) {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Href = removeTrailingSlash(attr.Val)
		}
	}

	return
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
