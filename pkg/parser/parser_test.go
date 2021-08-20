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

		version, err := FindHTMLVersion(strings.NewReader(example))
		checkError(t, err)

		assertHTMLVersion(t, "5.0", version)
	})

	t.Run("it returns an empty string if the HTML tag doesn't have a version attribute", func(t *testing.T) {
		example := `<html>
					<head></head>
					<body></body>
					</html>`

		version, err := FindHTMLVersion(strings.NewReader(example))
		checkError(t, err)

		assertHTMLVersion(t, "", version)
	})
}

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
	t.Run("it returns 1 as the amount of external links in the page", func(t *testing.T) {

		// This is the URL of the page our user wants to find details about
		pageURL := "https://example.com"
		example := `<html>
						<body>
							<a href="https://google.com">Google</a>
						</body>
					</html>`

		linksCount, err := FindAllLinks(strings.NewReader(example), pageURL)
		checkError(t, err)

		externalLinks := linksCount.External
		assertCount(t, 1, externalLinks)
	})

	t.Run("it returns 2 internal and 0 external links in the page", func(t *testing.T) {

		pageURL := "https://example.com"

		example := `<html>
						<body>
							<a href="https://example.com/about">About</a>
							<a href="https://blog.example.com">Blog</a>
						</body>
					</html>`

		linksCount, err := FindAllLinks(strings.NewReader(example), pageURL)
		checkError(t, err)

		internalLinks := linksCount.Internal
		externalLinks := linksCount.External

		assertCount(t, 2, internalLinks)
		assertCount(t, 0, externalLinks)
	})

	t.Run("it returns 1 internal and 1 external links in the page", func(t *testing.T) {

		pageURL := "https://example.com"

		example := `<html>
						<body>
							<a href="https://example.com/about">About</a>
							<a href="https://google.com">Google</a>
						</body>
					</html>`

		linksCount, err := FindAllLinks(strings.NewReader(example), pageURL)
		checkError(t, err)

		internalLinks := linksCount.Internal
		externalLinks := linksCount.External

		assertCount(t, 1, internalLinks)
		assertCount(t, 1, externalLinks)
	})

	t.Run("it returns only the count of unique external links", func(t *testing.T) {
		pageURL := "https://example.com"

		example := `<html>
						<body>
							<a href="https://google.com/about">About</a>
							<a href="https://google.com">Google</a>
							<a href="https://google.com/">Google</a>
						</body>
					</html>`

		links, err := FindAllLinks(strings.NewReader(example), pageURL)
		checkError(t, err)

		internalLinks := links.Internal
		externalLinks := links.External

		assertCount(t, 0, internalLinks)
		assertCount(t, 2, externalLinks)
	})

	t.Run("it returns 2 as the number of inaccessible links", func(t *testing.T) {
		pageURL := "https://example.com"
		example := `<html>
						<body>
							<a href="https://google.com">Google</a>
							<a href="https://somedomain.com/about">About</a>
							<a href="https://anotherdomain.com">Home</a>
						</body>
					</html>`

		links, err := FindAllLinks(strings.NewReader(example), pageURL)
		checkError(t, err)

		assertCount(t, 2, links.InAccessible)
	})

	t.Run("it returns 1 as the count of unique inaccessible links", func(t *testing.T) {
		pageURL := "https://example.com"
		example := `<html>
						<body>
							<a href="https://google.com">Google</a>
							<a href="https://somedomain.com">Some Domain</a>
							<a href="https://somedomain.com/">Some Domain</a>
						</body>
					</html>`

		links, err := FindAllLinks(strings.NewReader(example), pageURL)
		checkError(t, err)

		assertCount(t, 1, links.InAccessible)
	})
}

func TestCheckIfPageHasLoginForm(t *testing.T) {
	t.Run("it returns true if page contains a Login Form", func(t *testing.T) {
		example := `<form action="action_page.php" method="post">
						<label for="uname"><b>Username</b></label>
						<input type="text" placeholder="Enter Username" name="uname" required>
					
						<label for="psw"><b>Password</b></label>
						<input type="password" placeholder="Enter Password" name="psw" required>
					
						<button type="submit">Login</button>
					</form>`

		contains, err := CheckIfPageHasLoginForm(strings.NewReader(example))
		checkError(t, err)

		assertPageContainsLoginForm(t, true, contains)
	})

	t.Run("it returns false if page doesn't contain a Login Form", func(t *testing.T) {
		example := `<html>
					<title>Home Page</title>
					<body></body>
					</html`

		contains, err := CheckIfPageHasLoginForm(strings.NewReader(example))
		checkError(t, err)

		assertPageContainsLoginForm(t, false, contains)
	})
}

func assertHTMLVersion(t *testing.T, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected to find a HTML version value of %q, but got %q", expected, actual)
	}
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

func assertCount(t testing.TB, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected a count of %d, but got %d", expected, actual)
	}
}

func assertPageContainsLoginForm(t testing.TB, expected, actual bool) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected Login Form existence check to be %v, but got %v", expected, actual)
	}
}

func checkError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("test failed and encountered the following error: %v", err)
	}
}
