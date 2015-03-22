package main

import (
	"database/sql"
	"net/http"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/alert"
	"github.com/reinbach/zenpager/auth"
	"github.com/reinbach/zenpager/contacts"
	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/monitor"
	"github.com/reinbach/zenpager/session"
	"github.com/reinbach/zenpager/template"
)

var (
	db *sql.DB
)

func Intro(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, append(templates, "intro.html"),
		template.NewContext())
}

func Dashboard(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, append(templates, "dashboard.html"),
		template.NewContext())
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, append(templates, "404.html"),
		template.NewContext())
}

// ContextMiddleware creates a new go.net/context and
// injects into the current goji context.
func ContextMiddleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var ctx = context.Background()
		ctx = database.NewContext(ctx, db)

		// add the context to the goji web context
		webctx.Set(c, ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func main() {
	goji.Handle("/alert/*", alert.Router())
	goji.Handle("/auth/*", auth.Router())
	goji.Handle("/monitor/*", monitor.Router())

	http.HandleFunc(template.STATIC_URL, template.StaticHandler)
	goji.Get("/dashboard/", Dashboard)
	goji.Get("/dashboard", http.RedirectHandler("/dashboard/", 301))
	goji.Get("/", Intro)
	goji.NotFound(NotFound)

	// API v1
	goji.Handle("/api/v1/contacts/*", contacts.Routes())
	goji.Get("/api/v1/contacts",
		http.RedirectHandler("/api/v1/contacts/", 301))

	db = database.Connect()
	goji.Use(ContextMiddleware)
	goji.Use(session.Middleware)

	go monitor.Monitor()
	goji.Serve()
}
