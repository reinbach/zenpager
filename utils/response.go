package utils

import (
	"net/http"
)

type Message struct {
	Type    string
	Content string
}

type Response struct {
	Result   string
	Messages []Message
	ID       int64
}

func BadRequestResponse(w http.ResponseWriter, msg string) {
	m := Message{
		Type:    "danger",
		Content: msg,
	}
	HttpResponse(w, "error", http.StatusBadRequest, m)
}

func NotFoundResponse(w http.ResponseWriter, msg string) {
	m := Message{
		Type:    "danger",
		Content: msg,
	}
	HttpResponse(w, "error", http.StatusNotFound, m)
}

func HttpResponse(w http.ResponseWriter, t string, s int, m Message) {
	res := Response{
		Result:   t,
		Messages: []Message{m},
	}
	EncodePayload(w, s, res)
}
