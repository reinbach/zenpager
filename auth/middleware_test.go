package auth

import (
	"net/http/httptest"
	"testing"
)

// Authrequired
func TestAuthRequired(t *testing.T) {
	w := httptest.NewRecorder()
	AuthRequired(w)
}
