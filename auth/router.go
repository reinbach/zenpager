package auth

import (
	"strings"

	"github.com/zenazn/goji/web"
)

var (
	ROUTE_PREFIX = "/auth"
)

func Route(path string) string {
	return strings.Join([]string{ROUTE_PREFIX, path}, "")
}

func Router() *web.Mux {
	mux := web.New()
	mux.Get(Route("/login/"), Login)
	mux.Post(Route("/login/"), Login)
	return mux
}
