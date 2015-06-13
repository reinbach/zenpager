package auth

import (
	"database/sql"
	"net/http"

	"code.google.com/p/go.net/context"
	webctx "github.com/goji/context"
	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
)

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("X-Access-Token")
		if auth == "" {
			AuthRequired(w)
			return
		}
		t := Token{Token: auth}

		ctx := webctx.FromC(*c)

		var db = ctx.Value(database.DB_KEY).(*sql.DB)
		if err := t.Get(db); err == false {
			AuthRequired(w)
			return
		}
		ctx = context.WithValue(ctx, "user", t.User)
		webctx.Set(c, ctx)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func AuthRequired(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("{\"message\": \"Unauthorized\"}"))
}
