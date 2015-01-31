package session

import (
	"net"
	"net/http"
	"strings"
	"time"
)

func CreateCookie(r *http.Request, name, value string) http.Cookie {
	host, _, _ := net.SplitHostPort(r.Host)
	raw := strings.Join([]string{name, value}, "")
	unparsed := []string{raw}
	expires := time.Now().AddDate(0, 0, 1)
	rawexpires := expires.Format(time.UnixDate)
	maxage := 86400
	cookie := http.Cookie{
		name,
		value,
		"/",
		host,
		expires,
		rawexpires,
		maxage,
		true, // Secure
		true, // HttpOnly
		raw,
		unparsed,
	}

	return cookie
}
