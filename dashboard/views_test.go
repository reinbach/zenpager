package dashboard

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zenazn/goji/web"
)

func TestHomeView(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	c := web.C{}
	HomeView(c, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("200 expected, got %v instead", w.Code)
	}

}
