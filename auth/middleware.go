package auth

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// check for user/cookie/session
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
