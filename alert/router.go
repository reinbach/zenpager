package alert

import (
	"strings"

	"github.com/zenazn/goji/web"
)

var (
	PATH_PREFIX = "/alert"
)

func Route(p string) string {
	return strings.Join([]string{PATH_PREFIX, p}, "")
}

func Router() *web.Mux {
	mux := web.New()
	mux.Get(Route("/"), List)
	return mux
}
