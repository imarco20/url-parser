package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

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

	defer func(body io.Reader) {
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
