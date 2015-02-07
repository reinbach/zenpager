package session

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

// Store session data in web context
func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c.Env[SESSION_KEY] = ReadCookie(r)
		c.Env[MESSAGES_KEY] = GetMessageSession(c)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
