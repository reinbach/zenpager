package settings

import (
	"strings"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/auth"
)

var (
	PATH_PREFIX = "/settings"
)

func Route(p string) string {
	return strings.Join([]string{PATH_PREFIX, p}, "")
}

func Router() *web.Mux {
	mux := web.New()
	mux.Use(auth.Middleware)
	mux.Get(Route("/"), CommandView)
	mux.Get(Route("/commands/"), CommandView)
	mux.Get(Route("/contacts/"), ContactView)
	mux.Get(Route("/servers/"), ServerView)
	mux.Get(Route("/timeperiods/"), TimePeriodView)
	return mux
}
