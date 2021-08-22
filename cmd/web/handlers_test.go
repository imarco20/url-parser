package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

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
		{"returns invalid link message when submitting invalid page link", "http://", http.StatusOK, [][]byte{[]byte("This field is invalid")}},
		{"returns field cannot be blank message when submitting empty link", "", http.StatusOK, [][]byte{[]byte("This field cannot be blank")}},
		{"returns link details", "http://testpage.com", http.StatusOK, [][]byte{[]byte("http://testpage.com"), []byte("1.0"), []byte("Test Page")}},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			form := url.Values{}
			form.Add("link", tt.link)

			code, _, body := testSrv.postForm(t, "/details", form)

			if code != http.StatusOK {
				t.Errorf("expected status code %d; but got %d", tt.expectedCode, code)
			}

			for _, expectedBody := range tt.expectedBodies {
				if !bytes.Contains(body, expectedBody) {
					t.Errorf("expected response body %q to contain %q", body, expectedBody)
				}
			}
		})
	}
}
