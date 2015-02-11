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
	"git.ironlabs.com/greg/zenpager/session"
	"git.ironlabs.com/greg/zenpager/template"
)

var (
	datasource = "postgres://postgres@localhost/zenpager?sslmode=disable"
	db         *sql.DB
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, "intro/home.html", template.NewContext())
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, "intro/404.html", template.NewContext())
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
	goji.Handle("/monitor/*", monitor.Router())
	goji.Handle("/alert/*", alert.Router())
	goji.Handle("/dashboard/*", dashboard.Router())
	goji.Handle("/auth/*", auth.Router())
	http.HandleFunc(template.STATIC_URL, template.StaticHandler)
	goji.Get("/", Home)
	goji.NotFound(NotFound)

	db = database.Connect(datasource)
	goji.Use(ContextMiddleware)
	goji.Use(session.Middleware)

	go monitor.Monitor()
	goji.Serve()
}
