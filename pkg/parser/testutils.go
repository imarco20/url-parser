package parser

import (
	"net/http"
	"testing"
)

func mockHttpGetter(url string) (*http.Response, error) {
	accessibleLinks := map[string]bool{
		"https://google.com":        true,
		"https://somedomain.com":    false,
		"https://anotherdomain.com": false,
	}

	response := http.Response{}

	if accessibleLinks[url] == true {
		response.StatusCode = http.StatusOK
	} else {
		response.StatusCode = http.StatusInternalServerError
	}

	return &response, nil
}

func assertHeadingCount(t testing.TB, expected, actual int, headingLevel string) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected to find %d %q headings, but found %d", expected, headingLevel, actual)
	}
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

func assertError(t testing.TB, expected, actual error) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected to receive error %q but got %q", expected.Error(), actual.Error())
	}
}
