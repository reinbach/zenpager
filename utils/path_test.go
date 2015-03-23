package utils

import (
	"testing"
)

func TestGetAbsDir(t *testing.T) {
	d := GetAbsDir()
	l := d[len(d)-len(PARENT_PACKAGE):]
	if l != PARENT_PACKAGE {
		t.Errorf("Path needs to end with parent package, got %v instead", l)
	}
}

func TestGetAbsDirAdd(t *testing.T) {
	e := "test"
	d := GetAbsDir(e)
	l := d[len(d)-len(e):]
	if l != e {
		t.Errorf("Path needs to end with %v, got %v instead", d, l)
	}
}
