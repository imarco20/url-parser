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

		title, _ := FindTitle(strings.NewReader(htmlPage))

		assertPageTitle(t, "Test Page", title)
	})

	t.Run("it returns New Page as the page title", func(t *testing.T) {
		htmlPage := `<html>
					<head>
						<title>New Page</title>
					</head>
					</html>`

		title, _ := FindTitle(strings.NewReader(htmlPage))

		assertPageTitle(t, "New Page", title)
	})

	t.Run(`it returns "" on parsing a document with missing title tag`, func(t *testing.T) {
		htmlPage := `<html>
					<head>
					</head>
					</html>`

		title, _ := FindTitle(strings.NewReader(htmlPage))

		assertPageTitle(t, "", title)
	})

	t.Run(`it returns an error on parsing a document with missing title tag`, func(t *testing.T) {
		htmlPage := `<html>
					<head>
					</head>
					</html>`

		title, err := FindTitle(strings.NewReader(htmlPage))
		assertError(t, ErrPageTitleNotFound, err)

		assertPageTitle(t, "", title)
	})
}
