package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	rr := httptest.NewRecorder()

	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a mock http handler
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// pass mock http handler to secureHeaders() middleware
	secureHeaders(next).ServeHTTP(rr, r)

	rs := rr.Result()

	// examine response headers
	frameOptions := rs.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("want %q; got %q", "deny", frameOptions)
	}

	xssProtection := rs.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("want %q; got %q", "1; mode=block", xssProtection)
	}

	// check if middleware has called the next handler in line
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %q; got %q", http.StatusOK, rs.StatusCode)
	}

	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want response body data to equal %q", "OK")
	}

}
