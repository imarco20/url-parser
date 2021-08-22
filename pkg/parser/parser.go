package parser

import (
	"fmt"
	"golang.org/x/net/html"
)

var ErrHTMLVersionNotFound = fmt.Errorf("HTML version not found")
var ErrPageTitleNotFound = fmt.Errorf("page title not found")

// HTMLNode is a type that embeds *html.Node
type HTMLNode struct {
	*html.Node
}

// getNodes returns all nodes of the tag parameter
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

// buildLink is a node method that returns a link object and populates
// it with the value of its href attribute
func (node *HTMLNode) buildLink(pageURL string) (link Link) {

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			// Case 1: It's a relative path and doesn't have a host
			uriHost, _ := getUrlHost(attr.Val)
			if uriHost == "" {
				link.Href = pageURL + "/" + removeTrailingSlash(attr.Val)
			}
			// Case 2: The url has a host
			link.Href = removeTrailingSlash(attr.Val)
		}
	}

	return link
}

// buildSubmitterElement returns a submitter element and populates it with the
// value of its type attribute and its text
func (node *HTMLNode) buildSubmitterElement() (element SubmitterElement) {
	for _, attr := range node.Attr {
		if attr.Key == "type" {
			element.Type = attr.Val
		}
	}

	element.Text = getTextFromNode(node.Node)
	return
}

// buildTitle returns a title object and populates it with its text
func (node *HTMLNode) buildTitle() (title Title) {
	title.Value = getTextFromNode(node.Node)
	return
}

// buildVersion returns a HTML tag along with its version
func (node *HTMLNode) buildVersion() (htmlTag HTMLTag) {
	for _, attr := range node.Attr {
		if attr.Key == "version" {
			htmlTag.Version = attr.Val
		}
	}
	return
}
