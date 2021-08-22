package parser

import (
	"golang.org/x/net/html"
	"io"
)

// SubmitterElement is a type that represents HTML elements whose
// type attribute is submit
type SubmitterElement struct {
	Type string
	Text string
}

// CheckIfPageHasLoginForm returns a boolean that represents whether the page has a submitter element
// whose text is login
func CheckIfPageHasLoginForm(body io.Reader) (bool, error) {
	document, err := html.Parse(body)
	if err != nil {
		return false, err
	}

	// In case our submitter is a Button
	var buttons []SubmitterElement
	submitterNodes := getNodes(document, "button")

	for _, node := range submitterNodes {
		buttons = append(buttons, node.buildSubmitterElement())
	}

	loginKeywords := []string{"login", "Login", "Log In", "log in", "Sign In", "sign in"}
	for _, button := range buttons {
		if button.Type == "submit" {
			for _, keyword := range loginKeywords {
				if keyword == button.Text {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
