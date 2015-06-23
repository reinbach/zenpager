package database

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	e := Encrypt("hello world")

	if e == "hello world" {
		t.Errorf("Expected string to be encrypted, it was not")
	}

	r := EncryptMatch(e, "hello world")
	if r != true {
		t.Errorf("Expected a match")
	}

	f := EncryptMatch("wrong", "hello world")
	if f == true {
		t.Errorf("Did not expect a match")
	}
}
