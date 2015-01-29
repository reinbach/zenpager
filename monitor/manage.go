package monitor

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Add(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Do something monitory")
}
