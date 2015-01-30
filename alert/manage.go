package alert

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func List(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List all alerts in place")
}
