package session

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

var (
	//TODO make hashKey, blockKey set in config file
	hashKey  = []byte("3abe23ea4caabd558499d9f54f5798e7")
	blockKey = []byte("fe7c14dfa57ff69b4e6a274686ebb71e")
	s        = securecookie.New(hashKey, blockKey)
	n        = "session"
)

func CreateCookie(r *http.Request, encoded string) *http.Cookie {
	expires := time.Now().AddDate(0, 0, 1)
	cookie := &http.Cookie{
		Name:    n,
		Value:   encoded,
		Expires: expires,
		Path:    "/",
		// MaxAge: 86400,
		// Secure: true,
		// HttpOnly: true,
	}
	host, _, _ := net.SplitHostPort(r.Host)
	if host != "localhost" {
		cookie.Domain = host
	}
	return cookie
}

func GetValue(r *http.Request, key string) (interface{}, error) {
	session := ReadCookieHandler(r)
	if value, prs := session[key]; prs == true {
		return value, nil
	}
	return nil, errors.New(fmt.Sprintf("%v does not exist", key))
}

func SetCookieHandler(w http.ResponseWriter, r *http.Request, k, v string) error {
	session := ReadCookieHandler(r)
	session[k] = v
	if encoded, err := s.Encode(n, session); err == nil {
		cookie := CreateCookie(r, encoded)
		http.SetCookie(w, cookie)
	} else {
		return errors.New(
			fmt.Sprintf(
				"Cookie encoding issue: %v\nRandom Key: %x\n",
				err,
				securecookie.GenerateRandomKey(16),
			),
		)
	}
	return nil
}

func ReadCookieHandler(r *http.Request) map[string]string {
	session := make(map[string]string)
	if cookie, err := r.Cookie(n); err == nil {
		if err = s.Decode(n, cookie.Value, &session); err == nil {
			return session
		}
	}
	return session
}
