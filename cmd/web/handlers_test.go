package main

import (
	"bytes"
	"marcode.io/url-parser/pkg/forms"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	app := newTestApplication(t)

	t.Run("it renders the home page template successfully", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		app.homeHandler(response, request)

		var buf bytes.Buffer
		temp := app.templatesCache["home.page.tmpl"]
		err := temp.Execute(&buf, &templateData{Form: forms.New(nil)})
		if err != nil {
			t.Fatal(err)
		}

		assertResponseBody(t, buf.String(), response.Body.String())

		assertResponseCode(t, http.StatusOK, response.Code)
	})

	t.Run("it return method not allowed response status code when any HTTP verb other than GET is used", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()

		app.homeHandler(response, request)

		assertResponseCode(t, http.StatusMethodNotAllowed, response.Code)
	})
	t.Run("it return not found response status code when a request is sent to a target not specified in routes()", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/unavailable", nil)
		response := httptest.NewRecorder()

		app.homeHandler(response, request)

		assertResponseCode(t, http.StatusNotFound, response.Code)
	})
}

func TestShowDetailsHandler(t *testing.T) {
	app := newTestApplication(t)
	testSrv := newTestServer(app.routes())
	defer testSrv.Close()

	tests := []struct {
		testName       string
		link           string
		expectedCode   int
		expectedBodies [][]byte
	}{
		{"it returns invalid link message when submitting invalid page link", "http://", http.StatusOK, [][]byte{[]byte("This field is invalid")}},
		{"it returns field cannot be blank message when submitting empty link", "", http.StatusOK, [][]byte{[]byte("This field cannot be blank")}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			form := url.Values{}
			form.Add("link", tt.link)

			code, _, body := testSrv.postForm(t, "/details", form)

			assertResponseCode(t, http.StatusOK, code)

			for _, expectedBody := range tt.expectedBodies {
				if !bytes.Contains(body, expectedBody) {
					t.Errorf("expected response body %q to contain %q", body, expectedBody)
				}
			}
		})
	}

	t.Run("it renders details template successfully with given link details", func(t *testing.T) {

		form := url.Values{}
		form.Add("link", "https://testpage.com")

		code, _, body := testSrv.postForm(t, "/details", form)

		assertResponseCode(t, http.StatusOK, code)

		var buf bytes.Buffer
		tmpl := app.templatesCache["details.page.tmpl"]
		err := tmpl.Execute(&buf, &templateData{Link: mockGetLinkDetails("https://testpage.com")})
		if err != nil {
			t.Fatal(err)
		}

		assertResponseBody(t, buf.String(), string(body))
	})
	t.Run("it return method not allowed response status code when any HTTP verb other than POST is used", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/details", nil)
		response := httptest.NewRecorder()

		app.showDetailsHandler(response, request)

		assertResponseCode(t, http.StatusMethodNotAllowed, response.Code)
	})

}

func TestHealthCheckHandler(t *testing.T) {
	app := newTestApplication(t)

	t.Run("it returns a response with status 200 OK", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
		response := httptest.NewRecorder()

		app.healthCheckHandler(response, request)

		assertResponseCode(t, http.StatusOK, response.Code)

	})
}
