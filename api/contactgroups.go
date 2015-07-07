package api

import (
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
	"github.com/reinbach/zenpager/utils"
)

func ContactGroupList(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	d := models.ContactGroupGetAll(db)

	res := utils.Response{Result: "success", Data: d}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ContactGroupItem(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact Group not found.")
		return
	}

	group := models.ContactGroup{ID: id}
	group.Get(db)

	if group.Name == "" {
		utils.NotFoundResponse(w, "Contact Group not found")
		return
	}

	res := utils.Response{Result: "success", Data: group}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ContactGroupAdd(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	g := models.ContactGroup{}
	if err := utils.DecodePayload(r, &g); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}

	errors := g.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	gl := models.ContactGroupGetAll(db)
	f := false
	for _, v := range gl {
		if v.Name == g.Name {
			f = true
		}
	}
	if f {
		utils.BadRequestResponse(
			w,
			"Contact Group with this name already exists.",
		)
		return
	}
	if g.Create(db) {
		m := utils.Message{
			Type:    "success",
			Content: "Contact Group data added.",
		}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update Contact Group.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func ContactGroupUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	// get id of contact group to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact Group not found.")
		return
	}

	// set contact with current data
	g := models.ContactGroup{ID: id}
	g.Get(db)

	// update data with new data and ensure it is valid
	if err := utils.DecodePayload(r, &g); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	errors := g.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if g.Update(db) {
		m := utils.Message{
			Type:    "success",
			Content: "Contact Group data updated.",
		}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update Contact Group.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func ContactGroupDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact Group not found.")
		return
	}

	g := models.ContactGroup{ID: id}
	g.Delete(db)

	res := utils.Response{Result: "success"}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ContactGroupContacts(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact Group not found.")
		return
	}

	g := models.ContactGroup{ID: id}
	g.Get(db)

	if g.Name == "" {
		utils.NotFoundResponse(w, "Contact Group not found")
		return
	}

	g.GetContacts(db)

	res := utils.Response{Result: "success", Data: g}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ContactGroupAddContact(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact Group not found.")
		return
	}

	g := models.ContactGroup{ID: id}
	g.Get(db)

	if g.Name == "" {
		utils.NotFoundResponse(w, "Contact Group not found")
		return
	}

	contact := models.Contact{}
	if err := utils.DecodePayload(r, &contact); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	contact.Get(db)
	if contact.Name == "" {
		utils.NotFoundResponse(w, "Contact not found")
		return
	}

	s := g.AddContact(db, &contact)
	if s == true {
		g.Contacts = append(g.Contacts, contact)
		res := utils.Response{Result: "success", Data: g}
		utils.EncodePayload(w, http.StatusOK, res)
	} else {
		utils.BadRequestResponse(w, "Failed to add contact to group.")
	}
}

func ContactGroupRemoveContact(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact Group not found.")
		return
	}

	cid, err := strconv.ParseInt(c.URLParams["cid"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Contact not found.")
		return
	}

	g := models.ContactGroup{ID: id}
	contact := models.Contact{ID: cid}

	g.Get(db)
	if g.Name == "" {
		utils.NotFoundResponse(w, "Contact Group not found")
		return
	}

	s := g.RemoveContact(db, &contact)
	if s == true {
		cs := []models.Contact{}
		for _, ct := range g.Contacts {
			if ct.ID != contact.ID {
				cs = append(cs, ct)
			}
		}
		g.Contacts = cs
		res := utils.Response{Result: "success", Data: g}
		utils.EncodePayload(w, http.StatusOK, res)
	} else {
		utils.BadRequestResponse(w, "Failed to remove contact from group.")
	}
}
