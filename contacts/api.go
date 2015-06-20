package contacts

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	var db = database.FromContext(c)

	d := GetAll(db)

	res := utils.Response{Result: "success", Data: d}
	utils.EncodePayload(w, http.StatusOK, res)
}

func Item(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(
		strings.Replace(r.URL.Path, "/", "", 1),
		10,
		64,
	)
	if err != nil {
		utils.BadRequestResponse(w, "Invalid contact data.")
		return
	}

	contact := Contact{ID: id}
	contact.Get(db)

	log.Println(contact)
	res := utils.Response{Result: "success", Data: contact}
	utils.EncodePayload(w, http.StatusOK, res)
}

func Add(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	type Data struct {
		Name  string
		Email string
	}
	d := Data{}
	if err := utils.DecodePayload(r, &d); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}

	contact := Contact{}
	contact.Name = d.Name
	contact.User.Email = d.Email

	errors := contact.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	contact.User.GetByEmail(db)
	if contact.User.ID == 0 {
		contact.User.Create(db)
	}
	contact.GetByUser(db)
	if contact.ID != 0 {
		utils.BadRequestResponse(
			w,
			"Contact already exists for this email address.",
		)
		return
	}
	if contact.Create(db) {
		m := utils.Message{Type: "success", Content: "Contact data added."}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update contact.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
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

	// get id of contact to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact not found.")
		return
	}

	// set contact with current data
	contact := Contact{
		ID: id,
	}
	contact.Get(db)

	// update data with new data and ensure it is valid
	if err := utils.DecodePayload(r, &contact); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	errors := contact.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if contact.Update(db) {
		m := utils.Message{Type: "success", Content: "Contact data updated."}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update contact.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func Delete(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(
		strings.Replace(r.URL.Path, "/", "", 1),
		10,
		64,
	)
	if err != nil {
		utils.BadRequestResponse(w, "Invalid contact data.")
		return
	}

	contact := Contact{ID: id}
	contact.Delete(db)

	res := utils.Response{Result: "success"}
	utils.EncodePayload(w, http.StatusOK, res)
}
