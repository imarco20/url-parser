package parser

import (
	"strings"
	"testing"
)

func TestFindAllLinks(t *testing.T) {
	t.Run("it returns 1 as the amount of external links in the page", func(t *testing.T) {

		// This is the URL of the page our user wants to find details about
		pageURL := "https://example.com"
		example := `<html>
						<body>
							<a href="https://google.com">Google</a>
						</body>
					</html>`

		linksCount, _ := FindAllLinks(mockHttpGetter, strings.NewReader(example), pageURL)

		externalLinks := linksCount.External
		assertCount(t, 1, externalLinks)
	})

	t.Run("it returns 2 internal (sub-domain) and 0 external links in the page", func(t *testing.T) {

		pageURL := "https://example.com"

		example := `<html>
						<body>
							<a href="https://about.example.com">About</a>
							<a href="https://blog.example.com">Blog</a>
						</body>
					</html>`

		linksCount, _ := FindAllLinks(mockHttpGetter, strings.NewReader(example), pageURL)

		internalLinks := linksCount.Internal
		externalLinks := linksCount.External

		assertCount(t, 2, internalLinks)
		assertCount(t, 0, externalLinks)
	})

	t.Run("it returns 2 internal (relative) and 0 external links in the page", func(t *testing.T) {

		pageURL := "https://example.com"

		example := `<html>
						<body>
							<a href="/about">About</a>
							<a href="/blog">Blog</a>
						</body>
					</html>`

		linksCount, _ := FindAllLinks(mockHttpGetter, strings.NewReader(example), pageURL)

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

		linksCount, _ := FindAllLinks(mockHttpGetter, strings.NewReader(example), pageURL)

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

		links, _ := FindAllLinks(mockHttpGetter, strings.NewReader(example), pageURL)

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

		links, _ := FindAllLinks(mockHttpGetter, strings.NewReader(example), pageURL)

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

		links, _ := FindAllLinks(mockHttpGetter, strings.NewReader(example), pageURL)

		assertCount(t, 1, links.InAccessible)
	})
}
