package monitor

import (
	"strings"

	"github.com/zenazn/goji/web"
)

var (
	ROUTE_PREFIX = "/monitor"
)

func Route(path string) string {
	return strings.Join([]string{ROUTE_PREFIX, path}, "")
}

func Router() *web.Mux {
	mux := web.New()
	mux.Get(Route("/"), List)
	return mux
}
