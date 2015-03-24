package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func DecodePayload(r *http.Request, o interface{}) error {
	if r.ContentLength < 1 {
		return errors.New("Invalid or empty payload")
	}
	raw := make([]byte, r.ContentLength)
	_, err := io.ReadFull(r.Body, raw)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to read payload: %v", err))
	}
	err = json.Unmarshal(raw, &o)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to parse payload: %v", err))
	}
	return nil
}

func EncodePayload(w http.ResponseWriter, s int, r interface{}) {
	w.WriteHeader(s)
	b, err := json.Marshal(r)
	if err != nil {
		log.Println("Failed to encode response: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`'{"message": "Yikes, something broke"}'`))
	}
	w.Write(b)
}
