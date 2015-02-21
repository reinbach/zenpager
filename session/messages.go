package session

import (
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

func GetMessageSession(c *web.C) []string {
	if session, exists := c.Env[SESSION_KEY]; exists == true {
		if messages, exists := session.(Session)[MESSAGES_KEY]; exists == true {
			return messages.([]string)
		}
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
	if err := SetCookie(w, r, MESSAGES_KEY, nil); err != nil {
		log.Println("Failed to delete messages: ", err)
	}
}

func AddMessage(c *web.C, w http.ResponseWriter, r *http.Request, m string) {
	messages := GetMessageContext(c)
	messages = append(messages, m)
	if _, exists := c.Env[MESSAGES_KEY]; exists == true {
		c.Env[MESSAGES_KEY] = messages
	} else {
		c.Env = map[interface{}]interface{}{
			MESSAGES_KEY: messages,
		}
	}
	if err := SetCookie(w, r, MESSAGES_KEY, messages); err != nil {
		log.Println("Failed to set message: ", err)
	}
}

func GetMessages(c *web.C, w http.ResponseWriter, r *http.Request) []string {
	messages := GetMessageContext(c)
	DeleteMessageContext(c, w, r)
	return messages
}
