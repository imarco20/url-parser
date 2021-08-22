package parser

import (
	"strings"
	"testing"
)

func TestFindHeadingsCount(t *testing.T) {
	t.Run("it returns 1 as the count of h1 elements", func(t *testing.T) {
		example := `<html>
						<body>
							<h1>This page contains only one h1 element</h1>
						</body>
					</html>
					`
		headings, _ := FindAllHeadings(strings.NewReader(example))

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
		headings, _ := FindAllHeadings(strings.NewReader(example))

		// We should find 1 H2 Element
		assertHeadingCount(t, 1, headings.HTwo, "H2")
		// We should find 3 H4 Elements
		assertHeadingCount(t, 3, headings.HFour, "H4")
	})

	t.Run("it returns 1 as the count for each heading level", func(t *testing.T) {
		example := `<html>
						<body>
							<h1>Heading One</h1>
							<h2>Heading Two</h2>
							<h3>Heading Three</h3>
							<h4>Heading Four</h4>
							<h5>Heading Five</h5>
							<h6>Heading Six</h6>
						</body>
					</html>
					`
		headings, _ := FindAllHeadings(strings.NewReader(example))

		assertHeadingCount(t, 1, headings.HOne, "H1")
		assertHeadingCount(t, 1, headings.HTwo, "H2")
		assertHeadingCount(t, 1, headings.HThree, "H3")
		assertHeadingCount(t, 1, headings.HFour, "H4")
		assertHeadingCount(t, 1, headings.HFive, "H5")
		assertHeadingCount(t, 1, headings.HSix, "H6")
	})

	t.Run("it returns 0 as the count for each heading level", func(t *testing.T) {
		example := `<html>
						<body>
						</body>
					</html>
					`
		headings, _ := FindAllHeadings(strings.NewReader(example))

		assertHeadingCount(t, 0, headings.HOne, "H1")
		assertHeadingCount(t, 0, headings.HTwo, "H2")
		assertHeadingCount(t, 0, headings.HThree, "H3")
		assertHeadingCount(t, 0, headings.HFour, "H4")
		assertHeadingCount(t, 0, headings.HFive, "H5")
		assertHeadingCount(t, 0, headings.HSix, "H6")
	})
}
