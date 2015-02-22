package form

import (
	"net/http"
	"strings"
	"testing"
)

func TestIsValidTrue(t *testing.T) {
	f := NewForm()
	valid := f.IsValid()
	if valid == false {
		t.Error("Expected form to be valid")
	}
}

func TestIsValidFalse(t *testing.T) {
	f := NewForm()
	f.AddError("email", "wrong")
	valid := f.IsValid()
	if valid == true {
		t.Error("Expected form to be invalid")
	}
}

func TestAddError(t *testing.T) {
	f := NewForm()
	if len(f.Errors) != 0 {
		t.Errorf("Did not expect any errors, got %v instead", f.Errors)
	}
	f.AddError("email", "wrong")
	if len(f.Errors) != 1 {
		t.Errorf("Expected 1 error, got %v instead", f.Errors)
	}
	f.AddError("password", "wrong")
	if len(f.Errors) != 2 {
		t.Errorf("Expected 2 errors, got %v instead", f.Errors)
	}
}

func TestAddField(t *testing.T) {
	f := NewForm()
	f.AddField(EmailField)
	if len(f.Fields) != 1 {
		t.Errorf("Expected 1 field, got %v instead", f.Fields)
	}
}

func TestGetValue(t *testing.T) {
	f := NewForm()
	f.AddField(EmailField)
	v := f.GetValue("email")
	if v != EmailField.Value {
		t.Errorf("Expected %v, got %v instead", EmailField.Value, v)
	}
}

func TestValidate(t *testing.T) {
	r, _ := http.NewRequest("POST", "/",
		strings.NewReader("email=test@example.com"))
	f := Validate(r, []Field{EmailField})
	if len(f.Fields) != 1 {
		t.Errorf("Expected 1 field, got %v instead", f.Fields)
	}
}

func TestValidateError(t *testing.T) {
	r, _ := http.NewRequest("POST", "/", nil)
	f := Validate(r, []Field{EmailField})
	if len(f.Fields) != 0 {
		t.Errorf("Expected no fields, got %v instead", f.Fields)
	}
	if f.Errors["all"] != "Issue processing form data." {
		t.Errorf("Expected 1 error, got %v instead", f.Errors)
	}
}
