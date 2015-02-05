package dashboard

import (
	"strings"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/auth"
)

var (
	PATH_PREFIX = "/dashboard"
)

func Route(p string) string {
	return strings.Join([]string{PATH_PREFIX, p}, "")
}

func Router() *web.Mux {
	mux := web.New()
	mux.Use(auth.Middleware)
	mux.Get(Route("/"), View)
	return mux
}
