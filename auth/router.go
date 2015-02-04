package auth

import (
	"strings"

	"github.com/zenazn/goji/web"
)

func Route(u, p string) string {
	return strings.Join([]string{u, p}, "")
}

func Router(u string) *web.Mux {
	mux := web.New()
	mux.Get(Route(u, "/login/"), Login)
	mux.Post(Route(u, "/login/"), Login)
	mux.Get(Route(u, "/logout/"), Logout)
	return mux
}
