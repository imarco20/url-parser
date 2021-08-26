package main

import (
	"io"
	"log"
	"marcode.io/url-parser/pkg/models"
	"marcode.io/url-parser/pkg/parser"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// mockGetLinkDetails is a function that mocks the behavior of a function
// that takes a link url and returns its details
func mockGetLinkDetails(url string) models.LinkDetails {
	if url != "https://testpage.com" {
		return models.LinkDetails{}
	}
	return models.LinkDetails{
		PageURL:      url,
		HTMLVersion:  "1.0",
		Title:        "Test Page",
		Headings:     parser.HeadingCount{1, 2, 3, 4, 5, 6},
		Links:        parser.LinkCount{7, 8, 9},
		HasLoginForm: false,
	}
}

// newTestApplication returns an instance of our application
// composing mocked dependencies
func newTestApplication(t *testing.T) *application {
	templateCache, err := cacheAllTemplates("./../../ui/html")
	if err != nil {
		t.Fatal(err)
	}
	return &application{
		logger:         log.New(io.Discard, "", 0),
		templatesCache: templateCache,
		parser:         mockGetLinkDetails,
	}
}

// testServer is a custom test server that embeds an instance of the server
// provided by httptest
type testServer struct {
	*httptest.Server
}

// newTestServer initializes and returns an instance our custom defined testServer
// using the specified handler
func newTestServer(handler http.Handler) *testServer {
	server := httptest.NewServer(handler)
	return &testServer{server}
}

// postForm sends POST requests to the test server
// This function takes a form parameter which can hold the Form data
// we want to send in the request body
func (server *testServer) postForm(t *testing.T, path string, form url.Values) (int, http.Header, []byte) {
	response, err := server.Client().PostForm(server.URL+path, form)
	if err != nil {
		t.Fatal(err)
	}

	defer func(body io.ReadCloser) {
		err := response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	return response.StatusCode, response.Header, body
}

func assertResponseBody(t testing.TB, expected, actual string) {
	t.Helper()
	if expected != actual {
		t.Errorf("template isn't rendered as expected")
	}
}

func assertResponseCode(t *testing.T, expected, actual int) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected to get a %d response status code, but got %d", expected, actual)
	}
}
