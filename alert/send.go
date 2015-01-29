package alert

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Send(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Harass someone with an alert")
}
