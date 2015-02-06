package auth

import (
	"strings"

	"github.com/zenazn/goji/web"
)

var (
	PATH_PREFIX = "/auth"
)

func Route(p string) string {
	return strings.Join([]string{PATH_PREFIX, p}, "")
}

func Router() *web.Mux {
	mux := web.New()
	mux.Get(Route("/login/"), Login)
	mux.Post(Route("/login/"), Authenticate)
	mux.Get(Route("/logout/"), Logout)
	return mux
}
