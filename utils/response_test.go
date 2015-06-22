package utils

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// bad request response
func TestBadRequestResponse(t *testing.T) {
	w := httptest.NewRecorder()

	BadRequestResponse(w, "Failed")

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"Expected status code %v, got %v",
			http.StatusBadRequest,
			w.Code,
		)
	}

	if strings.Contains(fmt.Sprintf("%v", w.Body), "Failed") != true {
		t.Errorf("Expected 'Failed' in body, got '%v'", w.Body)
	}
}

// not found response
func TestNotFoundResponse(t *testing.T) {
	w := httptest.NewRecorder()

	NotFoundResponse(w, "404")

	if w.Code != http.StatusNotFound {
		t.Errorf(
			"Expected status code %v, got %v",
			http.StatusNotFound,
			w.Code,
		)
	}

	if strings.Contains(fmt.Sprintf("%v", w.Body), "404") != true {
		t.Errorf("Expected '404' in body, got '%v'", w.Body)
	}
}

// http response
func TestHttpResponse(t *testing.T) {
	w := httptest.NewRecorder()

	HttpResponse(w, "success", http.StatusAccepted, Message{Content: "Normal"})

	if w.Code != http.StatusAccepted {
		t.Errorf(
			"Expected status code %v, got %v",
			http.StatusAccepted,
			w.Code,
		)
	}

	if strings.Contains(fmt.Sprintf("%v", w.Body), "success") != true {
		t.Errorf("Expected 'success' in body, got '%v'", w.Body)
	}

	if strings.Contains(fmt.Sprintf("%v", w.Body), "Normal") != true {
		t.Errorf("Expected 'Normal' in body, got '%v'", w.Body)
	}
}
