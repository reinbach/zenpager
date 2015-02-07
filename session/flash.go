package session

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

func GetMessageSession(c *web.C) []string {
	session := c.Env[SESSION_KEY].(Session)
	if messages, exists := session[MESSAGES_KEY]; exists == true {
		return messages.([]string)
	}
	return []string{}
}

func GetMessageContext(c *web.C) []string {
	if messages, exists := c.Env[MESSAGES_KEY]; exists == true {
		return messages.([]string)
	}
	return []string{}
}

func DeleteMessageContext(c *web.C, w http.ResponseWriter, r *http.Request) {
	delete(c.Env, MESSAGES_KEY)
	if err := SetCookie(w, r, MESSAGES_KEY, ""); err != nil {
		log.Println("Failed to delete flash messages: ", err)
	}
}

func AddMessage(c *web.C, w http.ResponseWriter, r *http.Request, m string) {
	messages := GetMessageContext(c)
	messages = append(messages, m)
	c.Env[MESSAGES_KEY] = messages
	if err := SetCookie(w, r, MESSAGES_KEY, messages); err != nil {
		log.Println("Failed to set flash message: ", err)
	}
}

func GetMessages(c *web.C, w http.ResponseWriter, r *http.Request) []string {
	messages := GetMessageContext(c)
	DeleteMessageContext(c, w, r)
	return messages
}
