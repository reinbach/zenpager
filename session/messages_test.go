package session

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func TestGetMessageSessionEmptySession(t *testing.T) {
	c := &web.C{}
	s := GetMessageSession(c)
	if len(s) != 0 {
		t.Errorf("Expected empty string, got %v instead", s)
	}
}

func TestGetMessageSessionEmptyMessage(t *testing.T) {
	session := Session{}
	c := &web.C{
		Env: map[interface{}]interface{}{
			SESSION_KEY: session,
		},
	}
	s := GetMessageSession(c)
	if len(s) != 0 {
		t.Errorf("Expected empty string, got %v instead", s)
	}
}

func TestGetMessageSession(t *testing.T) {
	session := Session{}
	msg := []string{"test"}
	session[MESSAGES_KEY] = msg
	c := &web.C{
		Env: map[interface{}]interface{}{
			SESSION_KEY:  session,
			MESSAGES_KEY: msg,
		},
	}
	s := GetMessageSession(c)
	if s[0] != msg[0] {
		t.Errorf("Expected %v string, got %v instead", msg, s)
	}
}

func TestGetMessageContext(t *testing.T) {
	msg := []string{"test"}
	c := &web.C{
		Env: map[interface{}]interface{}{
			MESSAGES_KEY: msg,
		},
	}
	s := GetMessageContext(c)
	if s[0] != msg[0] {
		t.Errorf("Expected %v string, got %v instead", msg, s)
	}
}

func TestGetMessageContextEmpty(t *testing.T) {
	c := &web.C{}
	s := GetMessageContext(c)
	if len(s) != 0 {
		t.Errorf("Expected empty string, got %v instead", s)
	}
}

func TestDeleteMessage(t *testing.T) {
	msg := []string{"test"}
	c := &web.C{
		Env: map[interface{}]interface{}{
			MESSAGES_KEY: msg,
		},
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	DeleteMessageContext(c, w, r)

	// make sure it is gone
	s := GetMessageContext(c)
	if len(s) != 0 {
		t.Errorf("Expected message to be deleted, got %v instead", s)
	}
}

func TestDeleteMessageEmpty(t *testing.T) {
	c := &web.C{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	DeleteMessageContext(c, w, r)
	s := GetMessageContext(c)
	if len(s) != 0 {
		t.Errorf("Expected message to be deleted, got %v instead", s)
	}
}

func TestAddMessage(t *testing.T) {
	msg := "test"
	c := &web.C{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	AddMessage(c, w, r, msg)

	s := GetMessageContext(c)
	if len(s) != 0 && s[0] != msg {
		t.Errorf("Expected message %v, got %v instead", msg, s)
	}
}

func TestAddMessageAddition(t *testing.T) {
	msg := []string{"test"}
	c := &web.C{
		Env: map[interface{}]interface{}{
			MESSAGES_KEY: msg,
		},
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	AddMessage(c, w, r, "more")

	s := GetMessageContext(c)
	if len(s) != 2 && s[1] != "more" {
		t.Errorf("Expected message %v, got %v instead", msg, s)
	}
}

func TestGetMessage(t *testing.T) {
	msg := []string{"test"}
	c := &web.C{
		Env: map[interface{}]interface{}{
			MESSAGES_KEY: msg,
		},
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "", nil)
	s := GetMessages(c, w, r)

	if len(s) != 0 && s[0] != msg[0] {
		t.Errorf("Expected messages %v, got %v instead", msg, s)
	}

	s = GetMessageContext(c)
	if len(s) != 0 {
		t.Errorf("Expected messages to be deleted, got %v instead", s)
	}
}
