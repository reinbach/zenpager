package auth

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/session"
)

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if user, err := session.GetValue(r, USER_KEY); err == nil {
			// set user c
			log.Printf("user: %v\n", user)
		} else {
			// redirect to login page
			http.Redirect(w, r, Route("/login/"), http.StatusFound)
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
