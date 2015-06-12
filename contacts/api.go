package contacts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"

	"github.com/reinbach/zenpager/auth"
	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/utils"
)

func Routes() *web.Mux {
	api := web.New()
	api.Use(middleware.SubRouter)
	api.Use(auth.Middleware)

	// contacts
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
		log.Println("Contact Item Error: ", err)
	}
	io.WriteString(w, string(j))
}

func Add(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact Add Error: ", err)
	}
	io.WriteString(w, string(j))
}

func Update(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact Update Error: ", err)
	}
	io.WriteString(w, string(j))
}

func PartialUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)
	var res = utils.Response{}
	var m utils.Message

	// get id of contact to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		m = utils.Message{Type: "danger", Content: "Contact not found."}
		res = utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusNotFound, res)
		return
	}

	// set contact with current data
	contact := Contact{
		ID: id,
	}
	contact.Get(db)

	// update data with new data and ensure it is valid
	if err := utils.DecodePayload(r, &contact); err != nil {
		m = utils.Message{
			Type:    "danger",
			Content: "Data appears to be invalid.",
		}
		res = utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}
	errors := contact.Validate()
	if len(errors) > 0 {
		res = utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if contact.Update(db) {
		m = utils.Message{Type: "success", Content: "Contact data updated."}
		res = utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m = utils.Message{Type: "danger", Content: "Failed to update contact."}
		res = utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	contact := Contact{}
	j, err := json.Marshal(contact)
	if err != nil {
		log.Println("Contact Delete Error: ", err)
	}
	io.WriteString(w, string(j))
}
