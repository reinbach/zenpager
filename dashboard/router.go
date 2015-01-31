package dashboard

import (
	"strings"

	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/auth"
)

var (
	ROUTE_PREFIX = "/dashboard"
)

func Route(path string) string {
	return strings.Join([]string{ROUTE_PREFIX, path}, "")
}

func Router() *web.Mux {
	mux := web.New()
	mux.Use(auth.Middleware)
	mux.Get(Route("/"), View)
	return mux
}
