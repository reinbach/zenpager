package dashboard

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func View(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dashboard Display")
}
