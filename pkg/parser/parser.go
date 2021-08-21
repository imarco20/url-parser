package parser

import "golang.org/x/net/html"

type HTMLNode struct {
	*html.Node
}

func getNodes(node *html.Node, tag string) []*HTMLNode {

	// Base Case
	if node.Type == html.ElementNode && node.Data == tag {
		return []*HTMLNode{{node}}
	}

	var nodes []*HTMLNode
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		nodes = append(nodes, getNodes(child, tag)...)
	}

	return nodes
}

func (node *HTMLNode) buildLink() (link Link) {

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			link.Href = removeTrailingSlash(attr.Val)
		}
	}

	return link
}

func (node *HTMLNode) buildSubmitterElement() (element SubmitterElement) {
	for _, attr := range node.Attr {
		if attr.Key == "type" {
			element.Type = attr.Val
		}
	}

	element.Text = getTextFromNode(node.Node)
	return
}

func (node *HTMLNode) buildTitle() (title Title) {
	title.Value = getTextFromNode(node.Node)
	return
}

func (node *HTMLNode) buildVersion() (htmlTag HTMLTag) {
	for _, attr := range node.Attr {
		if attr.Key == "version" {
			htmlTag.Version = attr.Val
		}
	}
	return
}
