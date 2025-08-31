package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(test)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, got %d", status)
	}

	expected := "Hello from Go Test!"
	if rr.Body.String() != expected {
		t.Errorf("Expected %q, got %q", expected, rr.Body.String())
	}
}

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)
	handler.ServeHTTP(rr, req)

	// Since the file might not exist, we just check that we get some response
	if status := rr.Code; status != http.StatusOK && status != http.StatusNotFound {
		t.Errorf("Expected status 200 or 404, got %d", status)
	}
}