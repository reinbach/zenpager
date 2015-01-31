package auth

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/session"
)

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// check for user/cookie/session
		// if no user, then set anonymous
		user, err := r.Cookie("user")
		if err != nil {
			fmt.Println("Noooo: ", err)
		} else {
			fmt.Println("user: ", user.Value)
		}
		cookie := session.CreateCookie(r, "user", "Anonymous")
		http.SetCookie(w, &cookie)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
