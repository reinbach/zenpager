package session

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

var (
	f = "flash"
)

func GetMessageContext(c *web.C) Session {
	if session, exists := c.Env[SESSION_KEY]; exists == true {
		return session.(Session)
	}
	return Session{}
}

func DeleteMessageContext(c *web.C, w http.ResponseWriter, r *http.Request) {
	if _, exists := c.Env[SESSION_KEY]; exists == true {
		session := c.Env[SESSION_KEY].(Session)
		delete(session, f)
		c.Env[SESSION_KEY] = session
	}
	if err := SetCookie(w, r, f, ""); err != nil {
		log.Println("Failed to delete flash messages: ", err)
	}

}

func AddMessage(c *web.C, w http.ResponseWriter, r *http.Request, m string) {
	v := []string{}
	session := GetMessageContext(c)
	if value, exists := session[f]; exists == true {
		v = append(value.([]string), m)
	} else {
		v = append(v, m)
	}
	if err := SetCookie(w, r, f, v); err != nil {
		log.Println("Failed to set flash message: ", err)
	}
}

func GetMessages(c *web.C, w http.ResponseWriter, r *http.Request) []string {
	var msg []string
	session := GetMessageContext(c)
	if v, exists := session[f]; exists == true {
		msg = v.([]string)
		DeleteMessageContext(c, w, r)
	}
	return msg
}
