package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {

	// create application struct and add mock loggers
	app := newTestApplication(t)

	// create new test server
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// make a GET request to /ping
	statusCode, _, body := ts.get(t, "/ping")

	// examine response status code
	if statusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, statusCode)
	}

	// examine response body
	if string(body) != "OK" {
		t.Errorf("want response body data to equal %q", "OK")
	}

}

func TestShowSnippet(t *testing.T) {
	// create application struct instance
	app := newTestApplication(t)

	// establish new test server
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// set up table driven tests
	tests := []struct {
		name                 string
		urlPath              string
		requiredStatusCode   int
		requiredResponseBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("Mock content")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusCode, _, body := ts.get(t, tt.urlPath)

			if statusCode != tt.requiredStatusCode {
				t.Errorf("want %d; got %d", tt.requiredStatusCode, statusCode)
			}

			if !bytes.Contains(body, tt.requiredResponseBody) {
				t.Errorf("want response body data to contain %q", tt.requiredResponseBody)
			}
		})
	}
}
