package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/alert"
	"git.ironlabs.com/greg/zenpager/dashboard"
	"git.ironlabs.com/greg/zenpager/monitor"
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Sweet Home")
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Found (404)")
}

func main() {
	goji.Get("/monitor/add/", monitor.Add)
	goji.Get("/alert/", alert.Send)
	goji.Get("/dashboard/", dashboard.View)
	goji.Get("/", Home)
	goji.NotFound(NotFound)
	goji.Serve()
}
