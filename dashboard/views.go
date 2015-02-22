package dashboard

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/template"
)

var (
	templates = []string{
		"base.html",
		"dashboard/base.html",
		"dashboard/side.html",
	}
)

func HomeView(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, append(templates, "dashboard/index.html"),
		template.NewContext())
}
