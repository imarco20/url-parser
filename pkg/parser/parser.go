package parser

import (
	"golang.org/x/net/html"
	"io"
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

func FindAllLinks(body io.Reader) (int, error) {
	document, err := html.Parse(body)
	if err != nil {
		return 0, err
	}

	linkNodes := getNodes(document, "a")

	return len(linkNodes), nil
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
