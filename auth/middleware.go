package auth

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/session"
)

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if user, err := session.GetValue(r, "user"); err == nil {
			// set user c
			log.Println("user: %v", user)
		} else {
			// redirect to login page
			http.Redirect(w, r, "/auth/login/", http.StatusFound)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
