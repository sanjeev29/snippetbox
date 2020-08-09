package main

import (
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
