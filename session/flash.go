package session

import (
	"net/http"
)

func AddFlash(w http.ResponseWriter, r *http.Request, m interface{}) {
	//set cookie for flash
}
