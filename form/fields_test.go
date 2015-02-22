package form

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

var (
	EmailField = Field{
		Name:       "email",
		Required:   true,
		Validators: []Validator{Email{}},
		Value:      "",
	}
)

func TestEmailValidateTrue(t *testing.T) {
	e := Email{}
	valid, err := e.Validate(&EmailField, "test@example.com")
	if err != "" {
		t.Errorf("Validate should return empty string, got %v instead", err)
	}
	if valid == false {
		t.Errorf("Expected email to be valid")
	}
}

func TestEmailValidateFalse(t *testing.T) {
	e := Email{}
	valid, err := e.Validate(&EmailField, "testexample.com")
	if err != "Require valid email address." {
		t.Errorf("Validate should return proper string, got '%v' instead", err)
	}
	if valid == true {
		t.Errorf("Expected email to be invalid")
	}
}

func TestFieldValidate(t *testing.T) {
	r, _ := http.NewRequest("POST", "http://www.google.com/search?q=foo",
		strings.NewReader("email=test@example.com"))
	r.Header.Set("Content-Type",
		"application/x-www-form-urlencoded; param=value")
	if e := r.PostFormValue("email"); e != "test@example.com" {
		t.Errorf("PostFormValue missing form value, got '%v' instead", e)
	}
	v, valid, msg := EmailField.Validate(r)
	if v != "test@example.com" {
		t.Errorf("Expected email, got '%v' instead", v)
	}
	if valid != true {
		t.Error("Expected to pass validation")
	}
	if msg != "" {
		t.Errorf("Did not expect any message back, got '%v' instead", msg)
	}
}

func TestFieldValidateFalse(t *testing.T) {
	val := "testexample.com"
	r, _ := http.NewRequest("POST", "http://www.google.com/search?q=foo",
		strings.NewReader(fmt.Sprintf("email=%s", val)))
	r.Header.Set("Content-Type",
		"application/x-www-form-urlencoded; param=value")
	if e := r.PostFormValue("email"); e != val {
		t.Errorf("PostFormValue missing form value, got '%v' instead", e)
	}
	v, valid, msg := EmailField.Validate(r)
	if v != val {
		t.Errorf("Expected '%s', got '%v' instead", val, v)
	}
	if valid != false {
		t.Error("Expected to fail validation")
	}
	if msg != "Require valid email address." {
		t.Errorf("Expected valid message, got '%v' instead", msg)
	}
}
