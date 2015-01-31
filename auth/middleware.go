package auth

import (
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/session"
)

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if user, err := session.GetValue(r, "user"); err == nil {
			if user != "Anonymous" {
				// go and check whether user is valid
			}
		} else {
			session.SetCookieHandler(w, r, "user", "Anonymous")
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
