package main

import (
	"database/sql"
	"net/http"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/alert"
	"git.ironlabs.com/greg/zenpager/auth"
	"git.ironlabs.com/greg/zenpager/dashboard"
	"git.ironlabs.com/greg/zenpager/database"
	"git.ironlabs.com/greg/zenpager/monitor"
	"git.ironlabs.com/greg/zenpager/template"
)

var (
	datasource = "postgres://postgres@localhost/zenpager"
	db         *sql.DB
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(w, "intro/home.html", template.Context{})
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(w, "intro/404.html", template.Context{})
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
	http.Handle("/monitor/", monitor.Router())
	http.Handle("/alert/", alert.Router())
	http.Handle("/dashboard/", dashboard.Router())
	http.Handle("/auth/", auth.Router())
	http.HandleFunc(template.STATIC_URL, template.StaticHandler)
	goji.Get("/", Home)
	goji.NotFound(NotFound)

	db = database.Connect(datasource)
	goji.Use(ContextMiddleware)

	go monitor.Monitor()
	goji.Serve()
}
