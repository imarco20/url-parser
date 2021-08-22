package parser

import (
	"strings"
	"testing"
)

func TestFindHTMLVersion(t *testing.T) {
	t.Run("it returns the value of HTML version attribute if it exists", func(t *testing.T) {
		example := `<html version="5.0">
					<head></head>
					<body></body>
					</html>`

		version, _ := FindHTMLVersion(strings.NewReader(example))

		assertHTMLVersion(t, "5.0", version)
	})

	t.Run("it returns an empty string if the HTML tag doesn't have a version attribute", func(t *testing.T) {
		example := `<html>
					<head></head>
					<body></body>
					</html>`

		version, _ := FindHTMLVersion(strings.NewReader(example))

		assertHTMLVersion(t, "", version)
	})

	t.Run("it returns an error if the HTML version is not found", func(t *testing.T) {
		example := `<html>
						<head></head>
						<body></body>
					</html>
					`

		version, err := FindHTMLVersion(strings.NewReader(example))
		assertError(t, ErrHTMLVersionNotFound, err)

		assertHTMLVersion(t, "", version)
	})

	t.Run("it returns an empty string if HTML file does not contain HTML tag", func(t *testing.T) {
		example := `<body></body>`

		version, err := FindHTMLVersion(strings.NewReader(example))
		assertError(t, ErrHTMLVersionNotFound, err)

		assertHTMLVersion(t, "", version)
	})
}
