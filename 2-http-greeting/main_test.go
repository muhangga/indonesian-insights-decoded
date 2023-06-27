package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreeting(t *testing.T) {
	t.Run("with name", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/hello?name=Angga", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %v, but got %v", http.StatusOK, rr.Code)
		}

		expectedContentType := "application/json"
		if rr.Header().Get("Content-Type") != expectedContentType {
			t.Errorf("Expected Content-Type header to be %s, but got %s", expectedContentType, rr.Header().Get("Content-Type"))
		}

		expectedResponse := `{"message":"Hello, Angga!"}`
		if rr.Body.String() != expectedResponse {
			t.Errorf("Expected response body to be %s, but got %s", expectedResponse, rr.Body.String())
		}
	})

	t.Run("without name", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/hello", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status code %v, but got %v", http.StatusOK, rr.Code)
		}

		expectedContentType := "application/json"
		if rr.Header().Get("Content-Type") != expectedContentType {
			t.Errorf("Expected Content-Type header to be %s, but got %s", expectedContentType, rr.Header().Get("Content-Type"))
		}

		expectedResponse := `{"message":"Hello, World!"}`
		if rr.Body.String() != expectedResponse {
			t.Errorf("Expected response body to be %s, but got %s", expectedResponse, rr.Body.String())
		}
	})
}
