package api

import (
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/reinbach/zenpager/database"
	"github.com/reinbach/zenpager/models"
	"github.com/reinbach/zenpager/utils"
)

func CommandList(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	d := models.CommandGetAll(db)

	res := utils.Response{Result: "success", Data: d}
	utils.EncodePayload(w, http.StatusOK, res)
}

func CommandItem(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command not found.")
		return
	}

	command := models.Command{ID: id}
	command.Get(db)

	if command.Name == "" {
		utils.NotFoundResponse(w, "Command not found")
		return
	}

	res := utils.Response{Result: "success", Data: command}
	utils.EncodePayload(w, http.StatusOK, res)
}

func CommandAdd(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	type Data struct {
		Name    string
		Command string
	}
	d := Data{}
	if err := utils.DecodePayload(r, &d); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}

	command := models.Command{}
	command.Name = d.Name
	command.Command = d.Command

	errors := command.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if command.Create(db) {
		m := utils.Message{Type: "success", Content: "Command data added."}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update command.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func CommandUpdate(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	// get id of command to be updated
	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command not found.")
		return
	}

	// set command with current data
	command := models.Command{
		ID: id,
	}
	command.Get(db)

	// update data with new data and ensure it is valid
	if err := utils.DecodePayload(r, &command); err != nil {
		utils.BadRequestResponse(w, "Data appears to be invalid.")
		return
	}
	errors := command.Validate()
	if len(errors) > 0 {
		res := utils.Response{Result: "error", Messages: errors}
		utils.EncodePayload(w, http.StatusBadRequest, res)
		return
	}

	if command.Update(db) {
		m := utils.Message{Type: "success", Content: "Command data updated."}
		res := utils.Response{Result: "success", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusAccepted, res)
	} else {
		m := utils.Message{
			Type:    "danger",
			Content: "Failed to update command.",
		}
		res := utils.Response{Result: "error", Messages: []utils.Message{m}}
		utils.EncodePayload(w, http.StatusInternalServerError, res)
	}
}

func CommandDelete(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command not found.")
		return
	}

	command := models.Command{ID: id}
	command.Delete(db)

	res := utils.Response{Result: "success"}
	utils.EncodePayload(w, http.StatusOK, res)
}

func CommandGroups(c web.C, w http.ResponseWriter, r *http.Request) {
	var db = database.FromContext(c)

	id, err := strconv.ParseInt(c.URLParams["id"], 10, 64)
	if err != nil {
		utils.NotFoundResponse(w, "Command not found.")
		return
	}

	command := models.Command{ID: id}
	command.Get(db)

	if command.Name == "" {
		utils.NotFoundResponse(w, "Command not found")
		return
	}

	command.GetGroups(db)

	res := utils.Response{Result: "success", Data: command}
	utils.EncodePayload(w, http.StatusOK, res)
}
