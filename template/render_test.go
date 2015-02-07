package template

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/zenazn/goji/web"
)

func TestCreateTemplateList(t *testing.T) {
	l := CreateTemplateList("test.html")
	if len(l) != 2 {
		t.Errorf("Expected list of 2 templates, got %v", len(l))
	}

	// test templates in subdir
	l = CreateTemplateList("intro/home.html")
	if len(l) != 3 {
		t.Errorf("Expected list of 3 templates, got %v", len(l))
	}
}

func TestStaticHandlerValid(t *testing.T) {
	p := path.Join(STATIC_URL, "js/bootstrap.min.js")
	r, _ := http.NewRequest("GET", p, nil)
	w := httptest.NewRecorder()
	StaticHandler(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("200 expected, got %v instead", w.Code)
	}
}

func TestStaticHandlerInValid(t *testing.T) {
	p := path.Join(STATIC_URL, "js/something.js")
	r, _ := http.NewRequest("GET", p, nil)
	w := httptest.NewRecorder()
	StaticHandler(w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("404 expected, got %v instead", w.Code)
	}
}

func TestRender(t *testing.T) {
	c := web.C{}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	Render(c, w, r, "intro/home.html", &Context{})
}
