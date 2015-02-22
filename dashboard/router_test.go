package dashboard

import (
	"testing"
)

func TestRoute(t *testing.T) {
	r := Route("/test")
	if r != "/dashboard/test" {
		t.Errorf("Expected route %v, got %v instead", "/dashboard/test", r)
	}
}

func TestRouter(t *testing.T) {
	r := Router()
	if r == nil {
		t.Error("Router returned nil")
	}
}
