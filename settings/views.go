package settings

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/template"
)

var (
	templates = []string{
		"base.html",
		"dashboard/base.html",
		"settings/side.html",
	}
)

func CommandView(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	ctx.Add("CommandPage", true)
	template.Render(c, w, r, append(templates, "settings/command.html"), ctx)
}

func ContactView(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	ctx.Add("ContactPage", true)
	template.Render(c, w, r, append(templates, "settings/contact.html"), ctx)
}

func ServerView(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	ctx.Add("ServerPage", true)
	template.Render(c, w, r, append(templates, "settings/server.html"), ctx)
}

func TimePeriodView(c web.C, w http.ResponseWriter, r *http.Request) {
	ctx := template.NewContext()
	ctx.Add("TimePeriodPage", true)
	template.Render(c, w, r, append(templates, "settings/timeperiod.html"),
		ctx)
}
