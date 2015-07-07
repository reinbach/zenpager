package api

import (
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
	"github.com/reinbach/zenpager/utils"
)

func ServerList(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	d := models.ServerGetAll(db)

	res := utils.Response{Result: "success", Data: d}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ServerItem(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server not found.")
		return
	}

	server := models.Server{ID: id}
	server.Get(db)

	if server.Name == "" {
		utils.NotFoundResponse(w, "Server not found")
		return
	}

	res := utils.Response{Result: "success", Data: server}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ServerAdd(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	type Data struct {
		Name string
		URL  string
	}
	d := Data{}
	if err := utils.DecodePayload(r, &d); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}

	server := models.Server{}
	server.Name = d.Name
	server.URL.Host = d.URL

	errors := server.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if server.Create(db) {
		m := utils.Message{Type: "success", Content: "Server data added."}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update server.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func ServerUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	// get id of server to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server not found.")
		return
	}

	// set server with current data
	server := models.Server{
		ID: id,
	}
	server.Get(db)

	// update data with new data and ensure it is valid
	if err := utils.DecodePayload(r, &server); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	errors := server.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if server.Update(db) {
		m := utils.Message{Type: "success", Content: "Server data updated."}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update server.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func ServerDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server not found.")
		return
	}

	server := models.Server{ID: id}
	server.Delete(db)

	res := utils.Response{Result: "success"}
	utils.EncodePayload(w, http.StatusOK, res)
}

func ServerGroups(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Server not found.")
		return
	}

	server := models.Server{ID: id}
	server.Get(db)

	if server.Name == "" {
		utils.NotFoundResponse(w, "Server not found")
		return
	}

	server.GetGroups(db)

	res := utils.Response{Result: "success", Data: server}
	utils.EncodePayload(w, http.StatusOK, res)
}
