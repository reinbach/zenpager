package session

import (
	"fmt"
	"log"
	"net/http"
)

var (
	f = "flash"
)

func AddFlash(w http.ResponseWriter, r *http.Request, m string) {
	session := ReadCookie(r)
	v := []string{}
	if value, prs := session[f]; prs == true {
		v = append(value.([]string), m)
	} else {
		v = append(v, m)
	}
	fmt.Println("msg to be flashed: ", v)
	if err := SetCookie(w, r, f, v); err != nil {
		log.Fatal("Failed to set add flash message: ", err)
	}
}

func GetFlash(w http.ResponseWriter, r *http.Request) []string {
	msg, err := GetValue(r, f)
	if err != nil {
		fmt.Println("msg is empty...")
		msg = []string{}
	}
	if err := DeleteCookie(w, r, f); err != nil {
		log.Fatal("Failed to clear flash messages: ", err)
	}
	return msg.([]string)
}
