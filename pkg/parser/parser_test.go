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
		checkError(t, err)

		assertPageTitle(t, "Test Page", title)
	})

	t.Run("it returns New Page as the page title", func(t *testing.T) {
		htmlPage := `<html>
					<head>
						<title>New Page</title>
					</head>
					</html>`

		title, err := FindTitle(strings.NewReader(htmlPage))
		checkError(t, err)

		assertPageTitle(t, "New Page", title)
	})

	t.Run(`it returns "" on parsing a document with missing title tag`, func(t *testing.T) {
		htmlPage := `<html>
					<head>
					</head>
					</html>`

		title, err := FindTitle(strings.NewReader(htmlPage))
		checkError(t, err)

		assertPageTitle(t, "", title)
	})
}

func TestFindHeadingsCount(t *testing.T) {
	t.Run("it returns 1 as the count of h1 elements", func(t *testing.T) {
		example := `<html>
						<body>
							<h1>This page contains only one h1 element</h1>
						</body>
					</html>
					`
		headings, err := FindAllHeadings(strings.NewReader(example))
		checkError(t, err)

		expected := 1
		actual := headings.HOne

		if expected != actual {
			t.Errorf("expected to find %d %q heading, but found %d", expected, "H1", actual)
		}
	})

	t.Run("it returns 1 h2 and 3 h4 elements as headings count", func(t *testing.T) {
		example := `<html>
						<body>
							<h2>This page contains only one h2 element</h2>
							<h4>And it also</h4>
							<h4>contains three</h4>
							<h4>h4 elements</h4>
						</body>
					</html>
					`
		headings, err := FindAllHeadings(strings.NewReader(example))
		checkError(t, err)

		// We should find 1 H2 Element
		assertHeadingCount(t, 1, headings.HTwo, "H2")
		// We should find 3 H4 Elements
		assertHeadingCount(t, 3, headings.HFour, "H4")
	})

	t.Run("it returns 0 as the count for each heading level", func(t *testing.T) {
		example := `<html>
						<body>
						</body>
					</html>
					`
		headings, err := FindAllHeadings(strings.NewReader(example))
		checkError(t, err)

		assertHeadingCount(t, 0, headings.HOne, "H1")
		assertHeadingCount(t, 0, headings.HTwo, "H2")
		assertHeadingCount(t, 0, headings.HThree, "H3")
		assertHeadingCount(t, 0, headings.HFour, "H4")
		assertHeadingCount(t, 0, headings.HFive, "H5")
		assertHeadingCount(t, 0, headings.HSix, "H6")
	})
}

func TestFindAllLinks(t *testing.T) {
	t.Run("it returns 1 as the amount of page links", func(t *testing.T) {

		example := `<html>
						<body>
							<a href="https://google.com">Google</a>
						</body>
					</html>`

		linksCount, err := FindAllLinks(strings.NewReader(example))
		checkError(t, err)

		expected := 1
		actual := linksCount

		if expected != actual {
			t.Errorf("expected to find %d link, but found %d", expected, actual)
		}
	})

	t.Run("it returns 0 as the amount of page links", func(t *testing.T) {

		example := `<html>
						<body>
						</body>
					</html>`

		linksCount, err := FindAllLinks(strings.NewReader(example))
		checkError(t, err)

		expected := 0
		actual := linksCount

		if expected != actual {
			t.Errorf("expected to find %d link, but found %d", expected, actual)
		}
	})
}

func assertPageTitle(t testing.TB, expected, actual string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected to get a title of %q, but got %q", expected, actual)
	}
}

func assertHeadingCount(t testing.TB, expected, actual int, headingLevel string) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected to find %d %q headings, but found %d", expected, headingLevel, actual)
	}
}

func checkError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("test failed and encountered the following error: %v", err)
	}
}
