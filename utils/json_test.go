package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type Sample struct {
	Message string
}

// decode payload
func TestDecodePayload(t *testing.T) {
	s := Sample{
		Message: "hello",
	}
	j, _ := json.Marshal(s)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("GET", "/random", b)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	ns := Sample{}
	err = DecodePayload(r, &ns)
	if err != nil {
		t.Errorf("Did not expect an error decoding payload, got %v", err)
	}

	if ns.Message != s.Message {
		t.Errorf("Expected data %v to match %v", s.Message, ns.Message)
	}
}

// decode empty payload
func TestDecodePayloadEmpty(t *testing.T) {
	r, err := http.NewRequest("POST", "/random", nil)
	s := Sample{}
	err = DecodePayload(r, &s)
	if err == nil {
		t.Errorf("Expected an error, got nothing")
	}

	if err.Error() != "Invalid or empty payload" {
		t.Errorf("Expected 'Invalid or empty payload', got %v", err.Error())
	}
}

// decode invalid payload
func TestDecodePayloadInvalid(t *testing.T) {
	s := Sample{
		Message: "hello",
	}
	j, _ := json.Marshal(s)
	b := bytes.NewBuffer(j)

	r, err := http.NewRequest("GET", "/random", b)
	if err != nil {
		t.Errorf("Unexpected error", err)
	}

	ns := []string{}
	err = DecodePayload(r, &ns)
	if err == nil {
		t.Errorf("Expected an error, got nothing")
	}

	if strings.Contains(err.Error(), "Failed to parse") != true {
		t.Errorf("Expected 'Failed to parse...', got %v", err.Error())
	}

}

// encode payload
func TestEncode(t *testing.T) {
	s := Sample{Message: "Hello World!"}
	w := httptest.NewRecorder()

	EncodePayload(w, http.StatusOK, s)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}
}
