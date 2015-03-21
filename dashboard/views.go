package dashboard

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/template"
)

var (
	templates = []string{}
)

func HomeView(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, append(templates, "dashboard.html"),
		template.NewContext())
}
