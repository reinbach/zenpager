package dashboard

import (
	"strings"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/auth"
)

func Route(u, p string) string {
	return strings.Join([]string{u, p}, "")
}

func Router(u string) *web.Mux {
	mux := web.New()
	mux.Use(auth.Middleware)
	mux.Get(Route(u, "/"), View)
	return mux
}
