package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/alert"
	"github.com/reinbach/zenpager/api"
	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/middleware"
	"github.com/reinbach/zenpager/monitor"
	"github.com/reinbach/zenpager/template"
	"github.com/reinbach/zenpager/utils"
)

var (
	db        *sql.DB
	templates = []string{}
)

func Dashboard(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(c, w, r, []string{"dashboard.html"}, template.NewContext())
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("%v/%v", utils.GetAbsDir(),
		"404.html"))
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
	goji.Handle("/monitor/*", monitor.Router())

	goji.Get("/", Dashboard)
	http.HandleFunc("/static/", template.StaticHandler)
	goji.NotFound(NotFound)

	// API v1
	goji.Handle("/api/v1/*", api.Routes("/api/v1"))

	db = database.Connect()
	goji.Use(ContextMiddleware)
	goji.Use(middleware.Session)

	go monitor.Monitor()
	goji.Serve()
}
