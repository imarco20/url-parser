package parser

import (
	"strings"
	"testing"
)

func TestFindTitle(t *testing.T) {
	t.Run("it returns Test Page as the page title", func(t *testing.T) {

		htmlPage := `<html>
					<head>
						<title>Test Page</title>
					</head>
					</html>`

		title, err := FindTitle(strings.NewReader(htmlPage))
		if err != nil {
			t.Fatalf("test failed and encountered the following error: %v", err)
		}

		expected := "Test Page"
		actual := title

		if actual != expected {
			t.Errorf("expected to get a title of %q, but got %q", expected, actual)
		}
	})

	t.Run("it returns New Page as the page title", func(t *testing.T) {
		htmlPage := `<html>
					<head>
						<title>New Page</title>
					</head>
					</html>`

		title, err := FindTitle(strings.NewReader(htmlPage))

		if err != nil {
			t.Fatalf("test failed and encountered the following error: %v", err)
		}

		expected := "New Page"
		actual := title

		if actual != expected {
			t.Errorf("expected to get a title of %q, but got %q", expected, actual)
		}
	})
}
