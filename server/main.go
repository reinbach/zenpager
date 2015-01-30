package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"

	"git.ironlabs.com/greg/zenpager/alert"
	"git.ironlabs.com/greg/zenpager/dashboard"
	"git.ironlabs.com/greg/zenpager/monitor"
	"git.ironlabs.com/greg/zenpager/template"
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	template.Render(w, "home.html", template.Context{Title: "Welcome!"})
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Not Found (404)")
}

func main() {
	http.Handle("/monitor/", monitor.Router())
	http.Handle("/alert/", alert.Router())
	http.Handle("/dashboard/", dashboard.Router())
	http.HandleFunc(template.STATIC_URL, template.StaticHandler)
	goji.Get("/", Home)
	goji.NotFound(NotFound)

	go monitor.Monitor()
	goji.Serve()
}
