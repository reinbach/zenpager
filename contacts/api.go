package contacts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/reinbach/zenpager/auth"
)

func Routes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(auth.Middleware)

	api.Get("/", List)
	api.Get("/:id", Item)
	api.Post("/", Add)
	api.Put("/:id", Update)
	api.Patch("/:id", PartialUpdate)
	api.Delete("/:id", Delete)

	return api
}

func List(c web.C, w http.ResponseWriter, r *http.Request) {
	log.Println("List contacts...")
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact List Error: ", err)
	}
	io.WriteString(w, string(j))
}

func Item(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact List Error: ", err)
	}
	io.WriteString(w, string(j))
}

func Add(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact List Error: ", err)
	}
	io.WriteString(w, string(j))
}

func Update(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact List Error: ", err)
	}
	io.WriteString(w, string(j))
}

func PartialUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact List Error: ", err)
	}
	io.WriteString(w, string(j))
}

func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact List Error: ", err)
	}
	io.WriteString(w, string(j))
}
