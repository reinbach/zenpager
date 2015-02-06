package dashboard

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/template"
)

func View(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(w, "dashboard/index.html", template.NewContext(&c, w, r))
}
