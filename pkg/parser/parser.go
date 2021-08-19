package parser

import (
	"golang.org/x/net/html"
	"io"
	"strings"
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
		return "", err
	}

	var title Title
	title = buildTitle(nodes[0])

	return title.Value, nil
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
