package session

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCookie(t *testing.T) {
	v := "encoded"
	r, _ := http.NewRequest("GET", "", nil)
	c := CreateCookie(r, v)
	if c.Name != n {
		t.Errorf("Expected session name to be %v, got %v", n, c.Name)
	}
	if c.Value != v {
		t.Errorf("Expected session value to be %v, got %v", v, c.Value)
	}
	if c.Path != "/" {
		t.Errorf("Expected session path to be /, got %v", c.Path)
	}
}

func TestSetCookie(t *testing.T) {
	k := "key"
	v := "value"
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	err := SetCookie(w, r, k, v)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
}

func TestReadCookie(t *testing.T) {
	session := make(map[string]interface{})
	k := "key"
	v := "value"
	session[k] = v
	r, _ := http.NewRequest("GET", "", nil)
	if encoded, err := s.Encode(n, session); err == nil {
		cookie := CreateCookie(r, encoded)
		r.AddCookie(cookie)
		c := ReadCookie(r)
		if c[k] != v {
			t.Errorf("Expected cookie back, got %v", c)
		}
	} else {
		t.Errorf("Cookie encoding failed, got %v", err)
	}
}

func TestGetValueValid(t *testing.T) {
	session := make(map[string]interface{})
	k := "key"
	v := "value"
	session[k] = v
	r, _ := http.NewRequest("GET", "", nil)
	if encoded, err := s.Encode(n, session); err == nil {
		cookie := CreateCookie(r, encoded)
		r.AddCookie(cookie)
		if v2, err := GetValue(r, k); err == nil {
			if v != v2 {
				t.Errorf("Expected value %v, got %v", v, v2)
			}
		} else {
			t.Errorf("Failed to get cookie value: %v", err)
		}
	} else {
		t.Errorf("Cookie encoding failed, got %v", err)
	}
}

func TestGetValueInValid(t *testing.T) {
	session := make(map[string]interface{})
	k := "key"
	v := "value"
	session[k] = v
	r, _ := http.NewRequest("GET", "", nil)
	if encoded, err := s.Encode(n, session); err == nil {
		cookie := CreateCookie(r, encoded)
		r.AddCookie(cookie)
		if v2, err := GetValue(r, "else"); err == nil {
			if v == v2 {
				t.Errorf("Did not expect value %v", v)
			}
		}
	} else {
		t.Errorf("Cookie encoding failed, got %v", err)
	}
}
