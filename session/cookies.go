package session

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"

	"github.com/reinbach/zenpager/config"
)

var (
	cfg      = config.GetConfig()
	hashKey  = []byte(cfg.SESSION_HASH_KEY)
	blockKey = []byte(cfg.SESSION_BLOCK_KEY)
	s        = securecookie.New(hashKey, blockKey)
)

type Session map[string]interface{}

func CreateCookie(r *http.Request, encoded string) *http.Cookie {
	expires := time.Now().AddDate(0, 0, 1)
	cookie := &http.Cookie{
		Name:    COOKIE_KEY,
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
	session := ReadCookie(r)
	if value, prs := session[key]; prs == true {
		return value, nil
	}
	return nil, errors.New(fmt.Sprintf("%v does not exist", key))
}

func SetCookie(w http.ResponseWriter, r *http.Request, k string, v interface{}) error {
	session := ReadCookie(r)
	if v == nil {
		delete(session, k)
	} else {
		session[k] = v
	}
	if encoded, err := s.Encode(COOKIE_KEY, session); err == nil {
		cookie := CreateCookie(r, encoded)
		http.SetCookie(w, cookie)
	} else {
		log.Println("SetCookie issue: ", err)
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

func DeleteCookie(w http.ResponseWriter, r *http.Request, k string) error {
	return SetCookie(w, r, k, nil)
}

func ReadCookie(r *http.Request) Session {
	session := Session{}
	if cookie, err := r.Cookie(COOKIE_KEY); err == nil {
		if err = s.Decode(COOKIE_KEY, cookie.Value, &session); err == nil {
			return session
		}
	}
	return session
}
