package auth

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

func Middleware(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("X-Access-Token")
		if auth == "" {
			AuthRequired(w)
			return
		}
		//TODO make sure access token is valid, otherwise respond 401
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func AuthRequired(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("{\"message\": \"Unauthorized\"}"))
}
