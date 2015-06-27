package middleware

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/session"
)

// Store session data in web context
func Session(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env[session.SESSION_KEY] = session.ReadCookie(r)
		c.Env[session.MESSAGES_KEY] = session.GetMessageSession(c)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
