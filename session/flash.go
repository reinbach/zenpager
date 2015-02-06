package session

import (
	"log"
	"net/http"
)

var (
	f = "flash"
)

func AddMessage(w http.ResponseWriter, r *http.Request, m string) {
	session := ReadCookie(r)
	v := []string{}
	if value, prs := session[f]; prs == true {
		v = append(value.([]string), m)
	} else {
		v = append(v, m)
	}
	if err := SetCookie(w, r, f, v); err != nil {
		log.Println("Failed to set flash message: ", err)
	}
}

func GetMessages(w http.ResponseWriter, r *http.Request) []string {
	msg, err := GetValue(r, f)
	if err != nil {
		msg = []string{}
	}
	if err := DeleteCookie(w, r, f); err != nil {
		log.Println("Failed to clear flash messages: ", err)
	}
	return msg.([]string)
}
