package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	// init new httptest.ResponseRecorder
	rr := httptest.NewRecorder()

	// init a new dummy http.Request
	r, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// call the ping handler function
	ping(rr, r)

	// get response from http.ResponseRecorder
	rs := rr.Result()

	// examine response status code
	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	}

	// validate response body
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("want response body data to equal %q", "OK")
	}

}
